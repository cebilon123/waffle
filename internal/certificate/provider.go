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

type LocalCertificatesProvider struct {
	customCaCerts [][]byte
	certPEMBlock  []byte
	keyPEMBlock   []byte
}

func NewLocalCertificatesProvider(caCerts [][]byte, certPEMBlock, keyPEMBlock []byte) *LocalCertificatesProvider {
	return &LocalCertificatesProvider{
		customCaCerts: caCerts,
		certPEMBlock:  certPEMBlock,
		keyPEMBlock:   keyPEMBlock,
	}
}

func (l *LocalCertificatesProvider) GetTLSCertificate() (*tls.Certificate, error) {
	cert, err := tls.X509KeyPair(l.certPEMBlock, l.keyPEMBlock)
	if err != nil {
		return nil, fmt.Errorf("load x509 key pair: %w", err)
	}

	return &cert, nil
}

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
