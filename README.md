# Solana gRPC
### An implementation of a gRPC server which forwards JSON RPC requests to a Solana RPC node.

```
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --plugin=protoc-gen-grpc-gateway=/Users/garrett.partenza/go/bin/protoc-gen-grpc-gateway proto/solana.proto
```
