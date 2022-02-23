# grpc-go-course
This repository is meant for keeping the notes and codes for gRPC tutorial in GoLang from Udemy.

## Setup
go get -u github.com/golang/protobuf/protoc-gen-go  
go get -u google.golang.org/grpc  
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  
brew install protobuf  

## Command to compile proto file
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.  

## Command to set GO111MODULE to auto
go env -w GO111MODULE=auto

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

# Server Streaming
## Output of server run
go run greet/greet_server/server.go  
Greet Server!  
GreetManyTimes function was invoked with greeting:{first_name:"Anand"  last_name:"Verma"}  

## Output of client run
go run greet/greet_client/client.go  
Hello..I am a client!  
Starting to do a Server streaming RPC...  
2022/02/23 11:11:29 Response from GreetManyTimes: HelloAnand number 0  
2022/02/23 11:11:30 Response from GreetManyTimes: HelloAnand number 1  
2022/02/23 11:11:31 Response from GreetManyTimes: HelloAnand number 2  
2022/02/23 11:11:32 Response from GreetManyTimes: HelloAnand number 3  
2022/02/23 11:11:33 Response from GreetManyTimes: HelloAnand number 4  
2022/02/23 11:11:34 Response from GreetManyTimes: HelloAnand number 5  
2022/02/23 11:11:35 Response from GreetManyTimes: HelloAnand number 6  
2022/02/23 11:11:36 Response from GreetManyTimes: HelloAnand number 7  
2022/02/23 11:11:37 Response from GreetManyTimes: HelloAnand number 8  
2022/02/23 11:11:38 Response from GreetManyTimes: HelloAnand number 9  

