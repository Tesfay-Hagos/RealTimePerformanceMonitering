package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var STAT_KEYS = []string{"MemTotal", "MemFree", "MemAvailable", "SwapTotal", "SwapCached", "SwapFree"}

func ReadMemoryStats() (map[string]int, error) {
	// /proc/meminfo contains memory info. Read the file and parse the fields we need
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stats := map[string]int{}
	for scanner.Scan() {
		key, value := ParseMemInfoLine(scanner.Text())
		if slices.Contains(STAT_KEYS, key) {
			stats[key] = value
		}
	}

	return stats, nil
}
func ParseMemInfoLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return res
}
