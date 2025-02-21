package workerpool

type Worker interface {
	SendRequest(request string) (string, error)
	Stop() error
}

func NewWorker(addr string) (nw *Worker, err error){

	return
}