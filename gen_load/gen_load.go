package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/andreyst/agents-test-bench/gen_load/producer"
)

// TODO: Fix package names
// TODO: Add metrics
// TODO: Add file output
// TODO: Add other TCP/RELP inputs

func main() {
	// TODO: Move test duration to flags
	deadline := time.Now().Add(time.Duration(10-1) * time.Second)
	waitGroup := new(sync.WaitGroup)

	// TODO: Move fluentbit duration to flags
	fluentBitTCPProducer, err := producer.NewFluentBitTCPProducer("localhost:9099", deadline, 100, waitGroup)
	if err != nil {
		fmt.Printf("Error initializing fluent bit producer: %s\n", err.Error())
		return
	}

	waitGroup.Add(1)
	go fluentBitTCPProducer.Produce()

	waitGroup.Wait()
}
