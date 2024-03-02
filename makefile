certs_windows:
	mkdir ".cert"
	openssl genrsa -out .cert/server.key 2048
	openssl ecparam -genkey -name secp384r1 -out .cert/server.key
	winpty openssl req -new -x509 -sha256 -key .cert/server.key -out .cert/server.crt -days 3650
