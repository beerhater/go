package server

import (
	"fmt"
	"io"
	"log"

	"github.com/haterbeer/metrics-agent/internal/storage"
	pb "github.com/haterbeer/metrics-agent/proto"
)

type MetricsServer struct {
	pb.UnimplementedMetricsServiceServer
	Store *storage.Container 
}

func (s *MetricsServer) SendMetrics(stream pb.MetricsService_SendMetricsServer) error {
	log.Println("new agent connected")

	for {
		metrics, err := stream.Recv()

		if err == io.EOF {
			log.Println("agent disconnected (EOF)")
			return stream.SendAndClose(&pb.Empty{})
		}

		if err != nil {
			log.Printf("error receiving metric: %v", err)
			return err
		}

		s.Store.Save(metrics)

		count := len(s.Store.GetAll())
		
		fmt.Printf("[RAM: %d items] Metric: MemFree %d, Load %.2f\n", 
			count, metrics.Memory.Free, metrics.Cpu.Load_1)
	}
}

