package graph

import (
	"github.com/takuya911/golang-api_sample/services/graphql/graph/generated"
	pb "github.com/takuya911/golang-api_sample/services/graphql/proto"
)

// Resolver struct
type resolver struct {
	sampleClient pb.SampleServiceClient
}

// NewResolver function
func NewResolver(s pb.SampleServiceClient) generated.ResolverRoot {
	return &resolver{
		sampleClient: s,
	}
}
