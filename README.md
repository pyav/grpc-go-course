# grpc-go-course
This repository is meant for keep the notes and codes for gRPC tutorial in GoLang from Udemy.

go get -u github.com/golang/protobuf/protoc-gen-go  
go get -u google.golang.org/grpc  
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  
brew install protobuf  

# Command to compile proto file
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.  

# Unary RPC
## Output of server run
Greet Server!  
Greet function was invoked with greeting:{first_name:"Anand" last_name:"Verma"}  

## Output of client run
Hello..I am a client!  
Starting to do a unary RPC...  
2021/09/15 20:38:49 Response from Greet service: Hello Anand  

