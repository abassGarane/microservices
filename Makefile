.DEFAULT_GOAL:= swagger
swagger_install:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger: swagger_install
	GO111MODULE=off swagger generate spec -o ./swagger.yml --scan-models 

