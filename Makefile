.PHONY: protos

# gRPC protobuf constructors

protos_worker:
	protoc --proto_path=services/worker/pb --go_out=plugins=grpc:services/worker/pb service.proto

protos_coord:
	protoc --proto_path=services/coordinator/pb --go_out=plugins=grpc:services/coordinator/pb service.proto

protos_updater:
	protoc --proto_path=services/updater/pb --go_out=plugins=grpc:services/updater/pb service.proto

install:
	go mod download

BINARY_NAME=main.out
build_worker:
	go build -o services/worker/${BINARY_NAME} services/worker/main.go

run_worker:
	./services/worker/${BINARY_NAME}