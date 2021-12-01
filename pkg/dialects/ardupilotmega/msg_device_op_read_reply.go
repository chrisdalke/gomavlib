//autogenerated:yes
//nolint:golint,misspell,govet,lll
package ardupilotmega

// Read registers reply.
type MessageDeviceOpReadReply struct {
	// Request ID - copied from request.
	RequestId uint32
	// 0 for success, anything else is failure code.
	Result uint8
	// Starting register.
	Regstart uint8
	// Count of bytes read.
	Count uint8
	// Reply data.
	Data [128]uint8
	// Bank number.
	Bank uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageDeviceOpReadReply) GetID() uint32 {
	return 11001
}