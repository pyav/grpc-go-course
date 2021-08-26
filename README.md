# grpc-go-course

go get -u google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u github.com/golang/protobuf/protoc-gen-go

brew install protobuf

protoc greet/greetpb/greet.proto --go_out=.             

