check-swagger:
	which swagger || (go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	swagger generate spec -o ./docs/swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger ./docs/swagger.yaml