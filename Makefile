.PHONY: generate-proto run-client run-server

generate-proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

run-client:
	go run cmd/client/main.go

run-server:
	go run cmd/server/main.go
