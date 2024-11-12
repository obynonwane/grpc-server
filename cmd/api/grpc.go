// this is the server file
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/obynonwane/logger-service/logs"
	"google.golang.org/grpc"
)

// create a server type
type LogServer struct {
	logs.UnimplementedLogServiceServer
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	log.Println(input, "the input")
	log.Println(input.Data, "I just logged the data")

	// return response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

// start listening to tcp connection
func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{})

	log.Printf("gRPC Server started on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
