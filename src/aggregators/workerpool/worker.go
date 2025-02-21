package workerpool

import (
	"bytes"
	"net"
	"sync"
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

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(w.conn)
	if err != nil {
		return
	}
	response = buf.String()

	return
}