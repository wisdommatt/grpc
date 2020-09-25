proto:
	protoc -I protos/ protos/currency.proto --go_out=protos/currency

proto2:
	protoc -I=protos --go_out=protos/addressbook protos/addressbook.proto 

proto3:
	protoc --go_out=. person.proto

proto4:
	protoc --go-grpc_out=chat --go_out=chat chat.proto