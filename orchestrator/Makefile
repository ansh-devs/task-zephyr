NAME=orchestrator

proto:
	@ protoc --go_out=protov3 --go-grpc_out=protov3 protov3/taskzephyr.proto

build:
	@ go build -o bin/${NAME}_linux_amd64 cmd/main.go

run: build
	@ ./bin/${NAME}_linux_amd64