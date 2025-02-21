package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/CCathlete/featherweight/src/aggregators/workerpool"
)


func main() {
	workerCount := 4
	defaultBasePort := 5000
	basePort := flag.Int("port", defaultBasePort, "The base port for the workers")

	// A worker manager spawns Python workers and estblishes persistent 
	// connections.
	wm, err := workerpool.NewWorkerManager(workerCount, *basePort, workerpool.PY_WORKER)
	if err != nil {
		log.Fatalf("Failed to create worker manager: %v", err)
	}
	defer func () {
		errs := wm.StopAll()
		if errs != nil {
			log.Fatalf("Failed to stop workers: %v", errs)
		}
	}()

	fmt.Println("All workers started with persistent connections.")

	var wg sync.WaitGroup
	requestCount := 10
	for i := range(requestCount) {
		wg.Add(1)
		go func(reqNum int) {
			defer wg.Done()
			worker := wm.GetWorkers()[reqNum % workerCount]
			request := fmt.Sprintf("Request %d", reqNum)
			response, err := worker.SendRequest(request)
			if err != nil {
				fmt.Printf("Error sending request %d: %v\n", reqNum, err)
			}
			fmt.Printf("Response for request %d: %s\n", reqNum, response)
		}(i)
	}
	wg.Wait()

}