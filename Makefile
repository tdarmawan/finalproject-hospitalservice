GO_WORKSPACE := /Users/thomasdarmawan/Documents/go/gowork/src/testing/pb
PB_FILE := ./testing/pb

protoc:
	protoc  -I./proto --proto_path=proto proto/*.proto  --go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)
	@echo  "Protoc compile done!"

# grpcgateway:
#     protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. proto/user.proto
# 	protoc -I ./proto --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative ./proto/user.proto 
# 	@echo  "Protoc compile done!"


	