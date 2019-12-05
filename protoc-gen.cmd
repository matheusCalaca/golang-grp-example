protoc --proto_path=app/interface/rpc/proto/pessoa --proto_path=third_party --go_out=plugins=grpc:app/interface/rpc/api/pessoa endereco-service.proto
protoc --proto_path=app/interface/rpc/proto/pessoa --proto_path=third_party --go_out=plugins=grpc:app/interface/rpc/api/pessoa pessoa-service.proto
