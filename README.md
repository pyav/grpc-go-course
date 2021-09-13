# grpc-go-course
This repository is meant for keep the notes and codes for gRPC tutorial in GoLang from Udemy.

go get -u github.com/golang/protobuf/protoc-gen-go

go get -u google.golang.org/grpc

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

brew install protobuf

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
