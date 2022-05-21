package zmqwrapper

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

type ZmqWrapper struct {
	context *zmq.Context
	socket  *zmq.Socket
}

type Msg struct {
	Text string `json:"text"`
}

const (
	REPLY   = zmq.REP
	REQUEST = zmq.REQ
)

func NewZmqWrapper(mode zmq.Type) (*ZmqWrapper, error) {
	z := new(ZmqWrapper)
	z.context, _ = zmq.NewContext()
	z.socket, _ = z.context.NewSocket(zmq.Type(mode))
	return z, nil
}

func (z *ZmqWrapper) Connect(url string) {
	fmt.Println("Connecting...")
	z.socket.Connect(url)
}

func (z *ZmqWrapper) Bind(url string) {
	fmt.Println("Binding...")
	z.socket.Bind(url)
}

func (z *ZmqWrapper) Send(data string) {
	fmt.Println("Sending...")
	z.socket.Send(data, 0)
}

func (z *ZmqWrapper) Receive() (string, error) {
	fmt.Println("Receiving...")
	return z.socket.Recv(0)
}
