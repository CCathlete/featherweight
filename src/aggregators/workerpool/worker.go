package workerpool

import (
	"bufio"
	"fmt"
	"net"
	"sync"

	"github.com/CCathlete/featherweight/src/entities"
)

type PyWorker struct {
	addr string
	conn net.Conn
	mu sync.Mutex
}

// A cunstructor that establishes a persistent connection to a Python worker 
// (server).
func NewPyWorker(addr string) (nw *PyWorker, err error){
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	nw = &PyWorker{
		addr: addr,
		conn: conn,
	}

	return
}

// Sends the request and waits for a respinse. 
// We're using a mutex since workers are aorking in parallel.
func (w *PyWorker) SendRequest(request string) (response string, err error){
	w.mu.Lock()
	defer w.mu.Unlock()

	_, err = w.conn.Write([]byte(request))
	if err != nil {
		return
	}

	// Using a buffered io reader instead of a bytes buffer to avoid asynchronous 
	// behaviour.
	reader := bufio.NewReader(w.conn)
	response, err = reader.ReadString('\n')
	fmt.Printf("Request from client: %s\n", request)
	fmt.Printf("Response from worker: %s\n", response)

	return
}

// Termianting the persistent connection with this worker.
func (w *PyWorker) Stop() (err error) {
	err = w.conn.Close()
	return
}


var _ entities.Worker = (*PyWorker)(nil)