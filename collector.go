package main

import (
	"fmt"
	"github.com/mike-dunton/go-get-them-stats/helpers"
)

// WorkQueue is  a buffered channel that we can send work requests to.
var WorkQueue = make(chan helpers.WorkRequest, 100)

//Collector starts the collection of stats
func Collector() {
	//What stats we collecting?
	var name = "rabbitmq"
	work := helpers.WorkRequest{AppName: name}

	// Push the work onto the queue.
	WorkQueue <- work
	fmt.Println("Work request queued")

	return
}
