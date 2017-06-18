package signalchannel

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var terminationSignals = []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL}

// WaitForTermination blocks the process while waiting for a signal to end the process.
func WaitForTermination() os.Signal {
	signal := <-newTerminationChannel()
	fmt.Printf("Caught signal: %v \n", signal)
	return signal
}

func newTerminationChannel() chan os.Signal {

	var signalChannel chan os.Signal = make(chan os.Signal)
	registerSignals(signalChannel, terminationSignals)

	return signalChannel
}

func registerSignals(channel chan os.Signal, signals []os.Signal) {

	if len(signals) == 0 {
		return
	}

	for i := 0; i < len(signals); i++ {
		signal.Notify(channel, signals[i])
	}
}
