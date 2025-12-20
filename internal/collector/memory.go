package collector

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	Total int
	Free  int
}

func GetMemory() (Memory, error) {
	var mem Memory

	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return mem, err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				value, err := strconv.Atoi(fields[1])
				if err != nil {
					return mem, fmt.Errorf("ошибка в парсенге memtotal с линии %q: %w", line, err)
				}
				mem.Total = value
			}
		}

		if strings.HasPrefix(line, "MemFree:") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				value, err := strconv.Atoi(fields[1])
				if err != nil {
					return mem, fmt.Errorf("ошибка в парсенге memfree с линии %q: %w", line, err)
				}
				mem.Free = value
			}
		}
	}

	return mem, nil
}
