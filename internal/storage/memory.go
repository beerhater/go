package storage

import ( 
    "sync" 
    pb "github.com/haterbeer/metrics-agent/proto"
) 

type Container struct {
    mu   sync.RWMutex
    data map[int64]*pb.Metric 
}

func NewContainer() *Container { 
    return &Container{ 
        data: make(map[int64]*pb.Metric),
    } 
}

func (c *Container) Save(m *pb.Metric) { 
    c.mu.Lock() 
    defer c.mu.Unlock()  
    
    c.data[m.Timestamp] = m
} 

func (c *Container) GetAll() []*pb.Metric { 
    c.mu.RLock() 
    defer c.mu.RUnlock() 

    result := make([]*pb.Metric, 0, len(c.data))
    for _, v := range c.data { 
        result = append(result, v)
    } 
    return result 
}
