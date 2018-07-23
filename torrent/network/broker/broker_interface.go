package broker

type Request struct {
	address string
	data    []byte
}

func NewBrokerRequest(address string, data []byte) Request {
	return Request{
		address,
		data,
	}
}

type Broker interface {
	SendRequest(request Request) ([]byte, error)
}
