package main

import (
	"net"

	pb "github.com/jimmy-xu/learn-go/grpc/protos"
	"github.com/jimmy-xu/learn-go/grpc/util"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)
const (
	grpcPort = ":50051"
)
type Server struct{}
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	//get client info
	p, ok := peer.FromContext(ctx)
	if !ok {
		logrus.Errorf("failed to get peer of client")
	}
	logrus.Printf("receive gRPC request: [%v] client:%v%v", in.Name, p.Addr.String(), util.LineBreak)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logrus.Printf("failed to listen: %v%v", err, util.LineBreak)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &Server{})
	reflection.Register(grpcServer)

	logrus.Printf("Start gRPC Server on port :%v%v", grpcPort, util.LineBreak)
	grpcServer.Serve(listen)
}