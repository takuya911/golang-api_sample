package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/takuya911/golang-api_sample/services/graphql/graph"
	"github.com/takuya911/golang-api_sample/services/graphql/graph/generated"
	pb "github.com/takuya911/golang-api_sample/services/graphql/proto"
	"google.golang.org/grpc"
)

func main() {
	graphqlPort := os.Getenv("GRAPHQL_SERVICE_PORT")

	ctx := context.Background()
	// User gRPC connect
	grpcServerName := os.Getenv("GPRC_SERVICE_NAME")
	grpcPort := os.Getenv("GPRC_SERVICE_PORT")
	conn, err := grpc.DialContext(ctx, grpcServerName+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	sampleC := pb.NewSampleServiceClient(conn)

	resolver := graph.NewResolver(sampleC)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlPort)
	log.Fatal(http.ListenAndServe(":"+graphqlPort, nil))
}
