package main

import (
	"log"
	"net"
	"os"

	"github.com/takuya911/golang-api_sample/services/grpc/infrastructure"
	"github.com/takuya911/golang-api_sample/services/grpc/interface/controller"
	"github.com/takuya911/golang-api_sample/services/grpc/interface/repository"
	pb "github.com/takuya911/golang-api_sample/services/grpc/proto"
	"github.com/takuya911/golang-api_sample/services/grpc/usecase/interactor"
)

func main() {

	port := os.Getenv("GRPC_SERVICE_PORT")
	if port == "" {
		log.Fatalf("have to set port")
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// sampleservice
	sampleRepository := repository.NewSampleRepository()
	sampleInteractor := interactor.NewSampleInteractor(sampleRepository)
	sampleController := controller.NewSampleController(sampleInteractor)

	// services regist
	server := infrastructure.NewGrpcServer(sampleController)
	pb.RegisterSampleServiceServer(server, sampleController)

	log.Printf("starting server at http://localhost:%s/", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
