rm -rf ./proto/go
protoc -I=./protobuf/ --go_out=./proto protobuf/*