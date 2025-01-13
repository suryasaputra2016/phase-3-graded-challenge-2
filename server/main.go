package main

import (
	"log"
	"net"

	"github.com/robfig/cron/v3"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/proto/pb"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/servers"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("User server failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBookServiceServer(grpcServer, &servers.Server{})

	reflection.Register(grpcServer)

	// cron job
	c := cron.New()
	c.Start()
	s := &utils.Server{}
	s.StartSchedulerJob()

	log.Println("User server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("User server failed to serve: %v", err)
	}
}
