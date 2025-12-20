package main

import (
	"fmt"
	"log"
	"time"
	"github.com/haterbeer/metrics-agent/internal/collector"
)

func main() {
	for {
		mem, err := collector.GetMemory()
		if err != nil {
			log.Println("mem error:", err) 
		} 
		cpu, err := collector.GetCPU()
		if err != nil { 
			log.Println("CPU error:", err)
		} 
		fmt.Printf("STATE: MEMFREE: %d KB | Load: %.2f %.2f %.2f\n", mem.Free, cpu.Load1, cpu.Load5, cpu.Load15)

		time.Sleep(5 * time.Second)
	}
}
