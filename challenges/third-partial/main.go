package main

import (
	 "github.com/ferratverlataia/dc-labs/challenges/third-partial/api"
	"github.com/ferratverlataia/dc-labs/challenges/third-partial/controller"
	"github.com/ferratverlataia/dc-labs/challenges/third-partial/scheduler"
	"log"

)

func main() {
	log.Println("Welcome to the Distributed and Parallel Image Processing System")

	// Start Controller
	go controller.Start()

	// Start Scheduler
	jobs := make(chan scheduler.Job)
	go scheduler.Start(jobs)
	// Send sample jobs


	// API
	// Here's where your API setup will be
	api.Start()


}
