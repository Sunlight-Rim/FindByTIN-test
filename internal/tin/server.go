package tin

import (
	"io/ioutil"
	"net/http"
	pb "test-rusprofile/internal/tin/pb"

	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = "12201"
)

/// SERVER DEFINITION

type TinServiceServer struct {
	pb.UnimplementedTinServiceServer
}

func Start() {
	grpcServer := grpc.NewServer()
	tinService := TinServiceServer{}
	pb.RegisterTinServiceServer(grpcServer, &tinService)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("gRPC: TIN service listening at %v", grpcPort)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
}

/// API METHOD (gRPC)

func (s *TinServiceServer) Get(ctx context.Context, in *pb.GetTinRequest) (*pb.GetTinResponse, error) {
	log.Printf("Recieved: %v", in.GetTin())
	resp, err := http.Get("https://www.rusprofile.ru/search?query=" + in.GetTin())
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)
	// Getting info from https://www.rusprofile.ru/search?query=3664069397
	return &pb.GetTinResponse{Tin: "123"}, nil
}
