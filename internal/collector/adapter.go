package collector 


import ( 
		"time"
		pb "github.com/haterbeer/metrics-agent/proto"
) 

func MetricsToProto(mem Memory, cpu CPUStats) *pb.Metric { 
		return &pb.Metric{ 
        Timestamp: time.Now().Unix(), 

				Memory: &pb.MemoryStats{ 
						Total: int64(mem.Total), 
						Free: int64(mem.Free), 
				}, 

				Cpu: &pb.CPUStats{ 
						Load_1: cpu.Load1,
						Load_5: cpu.Load5,
						Load_15: cpu.Load15,
				}, 
		} 
} 
