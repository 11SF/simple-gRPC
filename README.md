###How to generate Protobuf file `(file_name.pb.go)`
1. go to directory `/protobuf` &&
 run command : `protoc --go_out=plugins=grpc:../ <file_name.proto>`