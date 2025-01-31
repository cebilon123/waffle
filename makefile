IMAGE_NAME =? <USERNAME>/<IMAGE-NAME>:<IMAGE-TAG>

certs_windows:
	mkdir -p "./cmd/proxy/.cert"

	# Create CA (certificate authority)
	openssl ecparam -out ./cmd/proxy/.cert/ca.key -name prime256v1 -genkey
	openssl req -new -sha256 -key ./cmd/proxy/.cert/ca.key -out ./cmd/proxy/.cert/ca.csr
	openssl x509 -req -sha256 -days 3650 -in ./cmd/proxy/.cert/ca.csr -signkey ./cmd/proxy/.cert/ca.key -out ./cmd/proxy/.cert/ca.crt

	# Create server certificate
	openssl ecparam -out ./cmd/proxy/.cert/server.key -name prime256v1 -genkey
	openssl req -new -sha256 -key ./cmd/proxy/.cert/server.key -out ./cmd/proxy/.cert/server.csr
	openssl x509 -req -in ./cmd/proxy/.cert/server.csr -CA ./cmd/proxy/.cert/ca.crt -CAkey ./cmd/proxy/.cert/ca.key -CAcreateserial -out ./cmd/proxy/.cert/server.crt -days 3650 -sha256
	openssl x509 -in ./cmd/proxy/.cert/server.crt -text -noout

certs:
	mkdir -p "./cmd/proxy/.cert"

	# Create CA (certificate authority)
	openssl ecparam -out ./cmd/proxy/.cert/ca.key -name prime256v1 -genkey
	openssl req -new -sha256 -key ./cmd/proxy/.cert/ca.key -out ./cmd/proxy/.cert/ca.csr
	openssl x509 -req -sha256 -days 3650 -in ./cmd/proxy/.cert/ca.csr -signkey ./cmd/proxy/.cert/ca.key -out ./cmd/proxy/.cert/ca.crt

	# Create server certificate
	openssl ecparam -out ./cmd/proxy/.cert/server.key -name prime256v1 -genkey
	openssl req -new -sha256 -key ./cmd/proxy/.cert/server.key -out ./cmd/proxy/.cert/server.csr
	openssl x509 -req -in ./cmd/proxy/.cert/server.csr -CA ./cmd/proxy/.cert/ca.crt -CAkey ./cmd/proxy/.cert/ca.key -CAcreateserial -out ./cmd/proxy/.cert/server.crt -days 3650 -sha256
	openssl x509 -in ./cmd/proxy/.cert/server.crt -text -noout

mocks:
	mockery

docker-build:
	docker build -t ${IMAGE_NAME} -f .\build\Dockerfile .

docker-push:
	docker push -t ${IMAGE_NAME} -f .\build\Dockerfile .
