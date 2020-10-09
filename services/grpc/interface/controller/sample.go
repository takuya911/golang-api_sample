package controller

import (
	"context"

	"github.com/takuya911/golang-api_sample/services/grpc/interface/usecase"
	pb "github.com/takuya911/golang-api_sample/services/grpc/proto"
)

type sampleController struct {
	sampleInteractor usecase.SampleUsecase
}

// NewSampleController function
func NewSampleController(u usecase.SampleUsecase) pb.SampleServiceServer {
	return &sampleController{u}
}

func (c *sampleController) GetRecode(ctx context.Context, in *pb.GetRecodeForm) (*pb.Recode, error) {
	var r *pb.Recode
	return r, nil
}

func (c *sampleController) CreateRecode(ctx context.Context, in *pb.CreateRecodeForm) (*pb.Recode, error) {
	var r *pb.Recode
	return r, nil
}
