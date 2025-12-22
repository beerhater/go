package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/haterbeer/metrics-agent/internal/collector"
	pb "github.com/haterbeer/metrics-agent/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("not connect", err)
	}
	defer conn.Close()

	client := pb.NewMetricsServiceClient(conn)

	stream, err := client.SendMetrics(context.Background())
	if err != nil {
		log.Fatal("cant open stream", err)
	}

	log.Println("Agent started. sending metrics to localhost:8080 ")

	for {
		mem, err := collector.GetMemory()
		if err != nil {
			log.Println("mem error", err)
			continue
		}

		cpu, err := collector.GetCPU()
		if err != nil {
			log.Println("cpu error", err)
			continue
		}

		metrics := collector.MetricsToProto(mem, cpu)

		if err := stream.Send(metrics); err != nil {
			log.Printf("failed to send metrics: %v", err)
		} else {
			fmt.Printf("send metrics: mem free %d, load %.2f\n", metrics.Memory.Free, metrics.Cpu.Load_1)
		}

		time.Sleep(5 * time.Second)
	}
}
