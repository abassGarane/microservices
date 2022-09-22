.PHONY:protos

protos:
	protoc -I protos/ --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. protos/currency.proto