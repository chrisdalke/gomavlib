//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Response from a PARAM_SET message when it is used in a transaction.
type MessageParamAckTransaction struct {
	// Id of system that sent PARAM_SET message.
	TargetSystem uint8
	// Id of system that sent PARAM_SET message.
	TargetComponent uint8
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter value (new value if PARAM_ACCEPTED, current value otherwise)
	ParamValue float32
	// Parameter type.
	ParamType MAV_PARAM_TYPE `mavenum:"uint8"`
	// Result code.
	ParamResult PARAM_ACK `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageParamAckTransaction) GetID() uint32 {
	return 19
}