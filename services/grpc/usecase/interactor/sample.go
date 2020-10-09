package interactor

import (
	"github.com/takuya911/golang-api_sample/services/grpc/usecase/repository"
)

type sampleInteractor struct {
	sampleRepository repository.SampleRepository
}

// NewSampleInteractor function
func NewSampleInteractor(u repository.SampleRepository) *sampleInteractor {
	return &sampleInteractor{u}
}

// SampleUsecase interface
type SampleUsecase interface {
}
