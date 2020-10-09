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

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	// User gRPC connect
	conn, err := grpc.DialContext(ctx, "grpc"+":"+"81", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	sampleC := pb.NewSampleServiceClient(conn)

	resolver := graph.NewResolver(sampleC)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
