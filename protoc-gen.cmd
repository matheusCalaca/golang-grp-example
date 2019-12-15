protoc --proto_path=app/interface/rpc/proto/pessoa --proto_path=third_party --go_out=plugins=grpc:app/interface/rpc/api/pessoa endereco-service.proto

protoc --proto_path=app/interface/rpc/proto/pessoa --proto_path=third_party --grpc-gateway_out=logtostderr=true:app/interface/rpc/api/pessoa --go_out=plugins=grpc:app/interface/rpc/api/pessoa --swagger_out=logtostderr=true:app/interface/rpc/api/pessoa pessoa-service.proto





