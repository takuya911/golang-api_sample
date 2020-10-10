package controller

import (
	"context"

	"github.com/golang/protobuf/ptypes"
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
	return &pb.Recode{
		Id:        1,
		Message:   "test",
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}, nil
}

func (c *sampleController) CreateRecode(ctx context.Context, in *pb.CreateRecodeForm) (*pb.Recode, error) {
	msg := in.GetMessage()

	return &pb.Recode{
		Id:        1,
		Message:   msg,
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}, nil
}
