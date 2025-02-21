package workerpool

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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

			exeDir := func() (ed string, err error){
				exePath, err := os.Executable()
				if err != nil {
					return
				}
				ed = filepath.Dir(exePath)

				return
			}

			if workerType == PY_WORKER {
				pwm := &PyWorkerManager{
					WorkerCount: workerCount,
					cmds: make([]*exec.Cmd, workerCount),
				}

				// Starting each of the workers.
				var ed, workerPath string
				for i := range workerCount {
					port := basePort + i
					// Spawning a py worker process.
					ed, err = exeDir()
					if err != nil {
						return
					}
					pyPath := fmt.Sprintf("%s/../../assets/python/worker.py", ed)
					workerPath, err = filepath.Abs(pyPath) 
					if err != nil {
						return
					}
					cmd := exec.Command("python3", workerPath, fmt.Sprintf("%d", port))
					if err = cmd.Start(); err != nil {
						return
					} 

					// Keeping track on the spawned processes.
					pwm.cmds = append(pwm.cmds, cmd)
					// Allowing the worker to start and bind to its port.
					time.Sleep(1500 * time.Millisecond)
					
					innerAddress := fmt.Sprintf("localhost:%d", port)
					var worker *PyWorker
					worker, err = NewPyWorker(innerAddress)
					if err != nil {
						return
					}
					pwm.Workers = append(pwm.Workers, worker)
				}
				nwm = pwm
			}

			return
		}