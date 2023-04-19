package gw

import (
	"context"
	"log"
	"net/http"
	pb "test-rusprofile/internal/tin/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	restPort = "8080"
)

/// SERVER DEFINITION

func Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterTinServiceHandlerFromEndpoint(ctx, rmux, "localhost:12201", opts); err != nil {
		panic(err)
	}
	// Handlers
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui"))))

	log.Printf("REST: Server listening at 8080")
	if err := http.ListenAndServe(":"+restPort, mux); err != nil {
		panic(err)
	}
}
