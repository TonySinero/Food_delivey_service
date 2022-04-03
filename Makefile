proto-order-ra:
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/order-ra.proto

proto-order-cs:
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/order-cs.proto

proto-user:
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/user.proto

proto-order-fd:
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/order-fd.proto

order-evans:
	evans api/proto/order.proto -p 50080

swag-generate:
	swag init -g cmd/main.go

run:
	go run cmd/main.go
