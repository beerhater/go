package main

import (
	"context"
	"fmt"
	"log"
	"os"      
	"time"

	"github.com/haterbeer/metrics-agent/internal/collector"
	pb "github.com/haterbeer/metrics-agent/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	target := os.Getenv("SERVER_ADDR")
	if target == "" {
		target = "localhost:8080"
	}

	log.Printf("Connecting to %s...", target)

	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", target, err)
	}
	defer conn.Close()

	client := pb.NewMetricsServiceClient(conn)

	stream, err := client.SendMetrics(context.Background())
	if err != nil {
		log.Fatalf("cant open stream: %v", err)
	}

	log.Printf("Agent started. Sending metrics to %s", target)

	for {
		mem, err := collector.GetMemory()
		if err != nil {
			log.Println("mem error:", err)
			continue
		}

		cpu, err := collector.GetCPU()
		if err != nil {
			log.Println("cpu error:", err)
			continue
		}

		metrics := collector.MetricsToProto(mem, cpu)

		if err := stream.Send(metrics); err != nil {
			log.Printf("failed to send metrics: %v", err)
		} else {
			fmt.Printf("sent metrics: mem free %d, load %.2f\n", metrics.Memory.Free, metrics.Cpu.Load_1)
		}

		time.Sleep(5 * time.Second)
	}
}
