package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func NewChartSettings(infile, outfile, width string) []string {
	result := []string{""}
	result = append(result, "-infile")
	result = append(result, infile)
	result = append(result, "-outfile")
	result = append(result, outfile)
	result = append(result, "-width")
	result = append(result, width)

	return result
}

func main() {
	chart1 := NewChartSettings("chartConfig.json", "chart.png", "850")
	runSysCall("/home/tudi/.config/versions/node/v12.6.0/bin/highcharts-export-server", chart1)
}

func runSysCall(command string, args []string) {
	binary, errPath := exec.LookPath(command)
	if errPath != nil {
		log.Println("look path error", errPath)
		os.Exit(99)
	}

	execErr := syscall.Exec(binary, args, os.Environ())
	if execErr != nil {
		log.Println("exec error", execErr)
		os.Exit(89)
	}
}
