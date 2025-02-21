package workerpool

import (
	"os/exec"

	"github.com/CCathlete/featherweight/src/entities"
)

type WorkerType string

const	PY_WORKER WorkerType = "py"


type WorkerManager interface {
	GetWorkers() []entities.Worker
	StopAll() []error
}

type PyWorkerManager struct {
	WorkerCount int
	Workers []entities.Worker
	// A list of commands that are running the workers. Useful for tracking.
	cmds []*exec.Cmd
}

func (pwm *PyWorkerManager) GetWorkers() []entities.Worker {
	return pwm.Workers
}

func (pwm *PyWorkerManager) StopAll() (errs []error) {
	// Stopping each worker.
	for _, worker := range pwm.Workers {
		worker.Stop()
	}
	// Killing the processes in the OS.
	for _, cmd := range pwm.cmds {
		err := cmd.Process.Kill()
		if err != nil {
			errs = append(errs, err)
		}
	}

	return
}

func 
	NewWorkerManager(
		workerCount int,
		basePort int,
		workerType WorkerType,
		) (
			nwm WorkerManager,
			err error,
		) {
			if workerType == PY_WORKER {
				nwm = &PyWorkerManager{
					WorkerCount: workerCount,
				}
			}

			return
		}