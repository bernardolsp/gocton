package communication

type Communicator struct {
}

func NewCommunicator() (*Communicator, error) {
	return &Communicator{}, nil
}

func (c *Communicator) SendMessage(queue string, message []byte) error {
	return nil
}

func (c *Communicator) ReceiveMessage(queue string) ([]byte, error) {
	return nil, nil
}
