package workerpool

type WorkerManager interface {
	GetWorkers() []Worker
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