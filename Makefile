install_swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger@latest


swagger: install_swagger
	swagger generate spec -o ./swagger.yaml --scan-models
