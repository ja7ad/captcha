gen-proto:
	@echo "generating proto..."
	@protoc --proto_path=proto --go_out=./proto --go_opt paths=source_relative --go-grpc_out=./proto --go-grpc_opt paths=source_relative  proto/*.proto
	@protoc --proto_path=_example/protobuf/echo --proto_path=proto --go_out=./_example/protobuf/echo --go_opt paths=source_relative --go-grpc_out=./_example/protobuf/echo --go-grpc_opt paths=source_relative  _example/protobuf/echo/*.proto