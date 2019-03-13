// +build ignore

package main

import (
	"fmt"
	"github.com/gswly/gomavlib"
	"github.com/gswly/gomavlib/dialects/ardupilotmega"
)

// this is an example struct that implements io.ReadWriteCloser.
// it does not read anything and prints what it receives.
// the only requisite is that Close() must release Read().
type CustomTransport struct {
	readChan chan []byte
}

func NewCustomTransport() *CustomTransport {
	return &CustomTransport{
		readChan: make(chan []byte),
	}
}

func (c *CustomTransport) Close() error {
	close(c.readChan)
	return nil
}

func (c *CustomTransport) Read(buf []byte) (int, error) {
	read,ok := <-c.readChan
	if ok == false {
		return 0, fmt.Errorf("terminated")
	}

	n := copy(buf,read)
	return n, nil
}

func (c *CustomTransport) Write(buf []byte) (int, error) {
	fmt.Println("sent:", buf)
	return len(buf), nil
}

func main() {
	// allocate the custom transport
	transport := NewCustomTransport()

	// create a node which understands given dialect, writes messages with given
	// system id and component id, and reads/writes through a custom transport.
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Dialect:     ardupilotmega.Dialect,
		SystemId:    10,
		ComponentId: 1,
		Transports: []gomavlib.TransportConf{
			gomavlib.TransportCustom{ transport },
		},
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	// queue a dummy message
	transport.readChan <- []byte("\xfd\t\x01\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x01\x02\x03\x05\x03\xd9\xd1\x01\x02\x00\x00\x00\x00\x00\x0eG\x04\x0c\xef\x9b")

	// work in a loop
	for {
		// wait until a message is received.
		res, ok := node.Read()
		if ok == false {
			break
		}

		// print message details
		fmt.Printf("received: id=%d, %+v\n", res.Message().GetId(), res.Message())
	}
}
