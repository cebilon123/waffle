certs_windows:
	mkdir ".cert"
	# Create CA (certificate authority)
	openssl ecparam -out .cert/ca.key -name prime256v1 -genkey
	openssl req -new -sha256 -key .cert/ca.key -out .cert/ca.csr
	openssl x509 -req -sha256 -days 3650 -in .cert/ca.csr -signkey .cert/ca.key -out .cert/ca.crt

	# Create server certificate
	openssl ecparam -out .cert/server.key -name prime256v1 -genkey
	openssl req -new -sha256 -key .cert/server.key -out .cert/server.csr
	openssl x509 -req -in .cert/server.csr -CA .cert/ca.crt -CAkey .cert/ca.key -CAcreateserial -out .cert/server.crt -days 3650 -sha256
	openssl x509 -in .cert/server.crt -text -noout

	# Create cert package in order to access .cert directory
	(echo package cert & echo "" & echo  "import (" & echo "\"embed\""  & echo " _ \"embed\"" & echo ")" & echo ""  & echo "//go:embed *" & echo "var Certificates embed.FS")  > .cert/cert.go
