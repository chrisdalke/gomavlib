//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/minimal"
)

// Version and capability of protocol version. This message can be requested with MAV_CMD_REQUEST_MESSAGE and is used as part of the handshaking to establish which MAVLink version should be used on the network. Every node should respond to a request for PROTOCOL_VERSION to enable the handshaking. Library implementers should consider adding this into the default decoding state machine to allow the protocol core to respond directly.
type MessageProtocolVersion = minimal.MessageProtocolVersion
