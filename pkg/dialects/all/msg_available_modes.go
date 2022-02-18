//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Get information about a particular flight modes.
// The message can be enumerated or requested for a particular mode using MAV_CMD_REQUEST_MESSAGE.
// Specify 0 in param2 to request that the message is emitted for all available modes or the specific index for just one mode.
// The modes must be available/settable for the current vehicle/frame type.
// Each modes should only be emitted once (even if it is both standard and custom).
type MessageAvailableModes struct {
	// The total number of available modes for the current vehicle type.
	NumberModes uint8
	// The current mode index within number_modes, indexed from 1.
	ModeIndex uint8
	// Standard mode.
	StandardMode MAV_STANDARD_MODE `mavenum:"uint8"`
	// System mode bitmap.
	BaseMode MAV_MODE_FLAG `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags
	CustomMode uint32
	// Name of custom mode, with null termination character. Should be omitted for standard modes.
	ModeName string `mavlen:"50"`
}

// GetID implements the msg.Message interface.
func (*MessageAvailableModes) GetID() uint32 {
	return 435
}
