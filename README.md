# grpc-go-course
This repository is meant for keeping the notes and codes for gRPC tutorial in GoLang from Udemy.

## Setup
go get -u github.com/golang/protobuf/protoc-gen-go  
go get -u google.golang.org/grpc  
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  
brew install protobuf  

## Command to compile proto file
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.  

# Unary RPC
## Output of server run
Greet Server!  
Greet function was invoked with greeting:{first_name:"Anand" last_name:"Verma"}  
## Output of client run
Hello..I am a client!  
Starting to do a unary RPC...  
2021/09/15 20:38:49 Response from Greet service: Hello Anand  
## Calculator service
### Output of server run
Calculator Server!  
Sum function was invoked with first_num:3  second_num:10  
### Output of client run
Calculator client  
Starting to do a Sum unary RPC...  
2021/09/15 22:41:10 Response from calculator service: 13  
