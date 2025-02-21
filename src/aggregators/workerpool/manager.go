package workerpool

import "github.com/CCathlete/featherweight/src/entities"

type WorkerManager interface {
	GetWorkers() []entities.Worker
	StopAll() error
}

func 
	NewWorkerManager(
		workerCount int,
		basePort int,
		) (
			nwm WorkerManager,
			err error,
		) {

			return
		}