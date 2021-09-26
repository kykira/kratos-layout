package generate

//go:generate protoc --proto_path=. --proto_path=./third_party --gogofaster_out=plugins=rpcx:. api/helloworld/v1/greeter.proto
