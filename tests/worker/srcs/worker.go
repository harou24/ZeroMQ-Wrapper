package main

import (
	"fmt"
	"strconv"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	fmt.Printf("Connecting to the server socket...\n")
	socket, _ := zctx.NewSocket(zmq.REP)
	socket.Connect("tcp://172.24.2.1:5555")
	fmt.Println("Connected")

	for {
		if msg, err := socket.Recv(0); err != nil {
			panic(err)
		} else {
			fmt.Println("Received msg->", msg)
			socket.Send(strconv.Itoa(len(msg)), 0)
		}
	}
}
