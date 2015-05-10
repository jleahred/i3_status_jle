package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	colorBlack   = "\x1b[30m"
	colorRed     = "\x1b[31m"
	colorGreen   = "\x1b[32m"
	colorYellow  = "\x1b[33m"
	colorBlue    = "\x1b[34m"
	colorMagenta = "\x1b[35m"
	colorCyan    = "\x1b[36m"
	colorWhite   = "\x1b[37m"
	colorDefault = "\x1b[39m"
)

func _ReadFileSplitLines(file string) []string {
	Txt, _ := ioutil.ReadFile(file)
	return strings.Split(string(Txt), "\n")
}

func _ParseMemInfoLine(memInfoLine string) (string, int64) {
	memName := strings.Trim(memInfoLine[0:17], " ")
	memName = memName[0 : len(memName)-1]
	memInt, _ := strconv.ParseInt(
		strings.Trim(memInfoLine[17:len(memInfoLine)-3], " "), 10, 0)
	return memName, memInt
}

func _MapFromMemInfoLines(lines []string) map[string]int64 {
	result := map[string]int64{}

	for _, line := range lines {
		if line != "" {
			name, value := _ParseMemInfoLine(line)
			result[name] = value
		}
	}

	return result
}

// Returns the current time
func getCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02  15:04:05")
}

func main() {
	hostname, _ := os.Hostname()

	for {
		memInfo := _MapFromMemInfoLines(
			_ReadFileSplitLines("/proc/meminfo"))
		fmt.Printf("%s |"+" RAM%%"+"  a: %.0f  swf: %.0f  f: %.0f  |  %s\n",
			hostname,
			//float64(100.*memInfo["MemAvailable"])/float64(memInfo["MemTotal"]),
			(float64(memInfo["MemFree"])+
				float64(memInfo["Cached"]))/
				float64(memInfo["MemTotal"])*100.,
			float64(100.*memInfo["SwapFree"])/float64(memInfo["SwapTotal"]),
			float64(100.*memInfo["MemFree"])/float64(memInfo["MemTotal"]),
			getCurrentTime())
		time.Sleep(3 * time.Second)
	}
}
