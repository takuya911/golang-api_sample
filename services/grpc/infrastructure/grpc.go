package infrastructure

import (
	pb "github.com/takuya911/golang-api_sample/services/grpc/proto"
	"google.golang.org/grpc"
)

// NewGrpcServer function
func NewGrpcServer(uc pb.SampleServiceServer) *grpc.Server {
	server := grpc.NewServer()
	return server
}
