package main

import (
	"log"
	"net"
	"os"         
	"os/signal"   
	"syscall"    

	"google.golang.org/grpc"

	"github.com/haterbeer/metrics-agent/internal/server"
	"github.com/haterbeer/metrics-agent/internal/storage"
	pb "github.com/haterbeer/metrics-agent/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	memStorage := storage.NewContainer()

	metricsHandler := &server.MetricsServer{
		Store: memStorage,
	}

	pb.RegisterMetricsServiceServer(grpcServer, metricsHandler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server is starting on :8080...")
		if err := grpcServer.Serve(listener); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-stop

	log.Println("\nShutting down gRPC server gracefully...")

	grpcServer.GracefulStop()

	log.Println("Server execution completed.")
}
