package entities

type Worker interface {
	SendRequest(request string) (string, error)
	Stop() error
}
