package main

import (
	"fmt"
	"os"

	"github.com/andreassiegel/timewatch/logger"
	"github.com/andreassiegel/timewatch/signalchannel"
	"github.com/andreassiegel/timewatch/timer"
)

func main() {

	fmt.Println("Starting timewatch...")

	logger.Initialize("log.csv")
	timer.Start()

	signalchannel.WaitForTermination()
	timer.Stop()

	logger.Add(timer.Summary().String())
	os.Exit(0)
}
