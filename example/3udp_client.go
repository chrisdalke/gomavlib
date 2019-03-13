// +build ignore

package main

import (
	"fmt"
	"github.com/gswly/gomavlib"
	"github.com/gswly/gomavlib/dialects/ardupilotmega"
)

func main() {
	// create a node which understands given dialect, writes messages with given
	// system id and component id, and reads/writes through udp in client mode.
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Dialect:     ardupilotmega.Dialect,
		SystemId:    10,
		ComponentId: 1,
		Transports: []gomavlib.TransportConf{
			gomavlib.TransportUdpClient{"1.2.3.4:5600"},
		},
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

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
