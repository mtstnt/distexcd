package main

import (
	"fmt"
	"log"

	"github.com/mtstnt/distexcd/pkg/core"
)

// Daemon App:
// 1. Start API and configures certificate (TODO).
// 2. Configure Docker client and versioning.
// 3. Configure resource management.
// 4. Reads from message queue (can we use internal queue for this?)

func main() {
	var initFunctions = []func() error{
		core.ConfigureLogger,
		core.ConfigureDockerClient,
	}

	fmt.Println("Running initialization functions...")
	for _, fn := range initFunctions {
		if err := fn(); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Hello world from Daemon!")
}
