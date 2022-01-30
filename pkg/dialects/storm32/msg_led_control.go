//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Control vehicle LEDs.
type MessageLedControl struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Instance (LED instance to control or 255 for all LEDs).
	Instance uint8
	// Pattern (see LED_PATTERN_ENUM).
	Pattern uint8
	// Custom Byte Length.
	CustomLen uint8
	// Custom Bytes.
	CustomBytes [24]uint8
}

// GetID implements the msg.Message interface.
func (*MessageLedControl) GetID() uint32 {
	return 186
}