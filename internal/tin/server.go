package tin

import (
	"errors"
	"strings"
	pb "test-rusprofile/internal/tin/pb"

	"context"
	"log"
	"net"

	"github.com/antchfx/htmlquery"

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
	// Get page HTML from https://www.rusprofile.ru/search?query=number
	doc, err := htmlquery.LoadURL("https://www.rusprofile.ru/search?query=" + in.GetTin())
	if err != nil {
		log.Fatalln(err)
	}
	// Parsing
	tgrc := htmlquery.FindOne(doc, "//*[@id='clip_kpp']")
	title := htmlquery.FindOne(doc, "//*[@id='ab-test-wrp']/div[1]/div[1]/h1")
	fcs := htmlquery.FindOne(doc, "//*[@id='anketa']/div[2]/div[1]/div[3]/span[3]/a/span")

	if tgrc != nil && title != nil && fcs != nil {
		return &pb.GetTinResponse{
			Tin:   in.GetTin(),
			Tgrc:  htmlquery.InnerText(tgrc),
			Title: strings.TrimSpace(strings.ReplaceAll(htmlquery.InnerText(title), `"`, "'")),
			FCs:   htmlquery.InnerText(fcs),
		}, nil
	} else {
		return nil, errors.New("parsing was wrong: TIN is correct and you haven't been banned from rusprofile.ru?")
	}
}
