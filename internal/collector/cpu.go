package collector 
import ( 
	"fmt"
	"os"
	"strconv"
	"strings"
)
type CPUStats struct { 
	Load1, Load5, Load15 float64 
} 
func GetCPU() (CPUStats, error) {
	var cpuStats CPUStats

	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return cpuStats, err
	}

	content := string(data)
	fields := strings.Fields(content)

	if len(fields) >= 3 {
		cpuStats.Load1, _ = strconv.ParseFloat(fields[0], 64)
		cpuStats.Load5, _ = strconv.ParseFloat(fields[1], 64)
		cpuStats.Load15, _ = strconv.ParseFloat(fields[2], 64)
		
		return cpuStats, nil 
	} 

	return cpuStats, fmt.Errorf("неверный формат данных: %s", content)
}
