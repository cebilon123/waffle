package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

// Provider should be implemented by the structs
// that are able to provide certificates.
type Provider interface {
	GetTLSCertificate() (*tls.Certificate, error)
	GetCACertificatesPool() (*x509.CertPool, error)
}

// LocalCertificatesProvider is a structure that holds custom CA certificates,
// the server certificate PEM block, and the private key PEM block.
// It provides methods for retrieving TLS certificates and a CA certificate pool for establishing secure connections.
type LocalCertificatesProvider struct {
	customCaCerts [][]byte
	certPEMBlock  []byte
	keyPEMBlock   []byte
}

// NewLocalCertificatesProvider initializes a new instance of LocalCertificatesProvider.
// It takes custom CA certificates, a certificate PEM block, and a key PEM block as input,
// and returns a LocalCertificatesProvider that can provide certificates for secure connections.
func NewLocalCertificatesProvider(caCerts [][]byte, certPEMBlock, keyPEMBlock []byte) *LocalCertificatesProvider {
	return &LocalCertificatesProvider{
		customCaCerts: caCerts,
		certPEMBlock:  certPEMBlock,
		keyPEMBlock:   keyPEMBlock,
	}
}

// GetTLSCertificate returns a TLS certificate created from the certificate PEM block and key PEM block.
// This is used to provide the server's certificate for TLS/SSL handshakes.
// If an error occurs while loading the key pair, it returns an error.
func (l *LocalCertificatesProvider) GetTLSCertificate() (*tls.Certificate, error) {
	cert, err := tls.X509KeyPair(l.certPEMBlock, l.keyPEMBlock)
	if err != nil {
		return nil, fmt.Errorf("load x509 key pair: %w", err)
	}

	return &cert, nil
}

// GetCACertificatesPool returns a pool of CA certificates, including both system CA certificates and
// any custom CA certificates provided to the LocalCertificatesProvider.
// If the system CA certificate pool cannot be loaded, it returns an error.
// Custom CA certificates are appended to the pool to support specific trusted CAs.
func (l *LocalCertificatesProvider) GetCACertificatesPool() (*x509.CertPool, error) {
	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("load system ca cert pool: %w", err)
	}

	if caCertPool == nil {
		caCertPool = x509.NewCertPool()
	}

	if len(l.customCaCerts) > 0 {
		for _, caCert := range l.customCaCerts {
			caCertPool.AppendCertsFromPEM(caCert)
		}
	}

	return caCertPool, nil
}
