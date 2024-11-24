//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Result from a MAVLink command (MAV_CMD)
type MAV_RESULT = common.MAV_RESULT

const (
	// Command is valid (is supported and has valid parameters), and was executed.
	MAV_RESULT_ACCEPTED MAV_RESULT = common.MAV_RESULT_ACCEPTED
	// Command is valid, but cannot be executed at this time. This is used to indicate a problem that should be fixed just by waiting (e.g. a state machine is busy, can't arm because have not got GPS lock, etc.). Retrying later should work.
	MAV_RESULT_TEMPORARILY_REJECTED MAV_RESULT = common.MAV_RESULT_TEMPORARILY_REJECTED
	// Command is invalid (is supported but has invalid parameters). Retrying same command and parameters will not work.
	MAV_RESULT_DENIED MAV_RESULT = common.MAV_RESULT_DENIED
	// Command is not supported (unknown).
	MAV_RESULT_UNSUPPORTED MAV_RESULT = common.MAV_RESULT_UNSUPPORTED
	// Command is valid, but execution has failed. This is used to indicate any non-temporary or unexpected problem, i.e. any problem that must be fixed before the command can succeed/be retried. For example, attempting to write a file when out of memory, attempting to arm when sensors are not calibrated, etc.
	MAV_RESULT_FAILED MAV_RESULT = common.MAV_RESULT_FAILED
	// Command is valid and is being executed. This will be followed by further progress updates, i.e. the component may send further COMMAND_ACK messages with result MAV_RESULT_IN_PROGRESS (at a rate decided by the implementation), and must terminate by sending a COMMAND_ACK message with final result of the operation. The COMMAND_ACK.progress field can be used to indicate the progress of the operation.
	MAV_RESULT_IN_PROGRESS MAV_RESULT = common.MAV_RESULT_IN_PROGRESS
	// Command has been cancelled (as a result of receiving a COMMAND_CANCEL message).
	MAV_RESULT_CANCELLED MAV_RESULT = common.MAV_RESULT_CANCELLED
	// Command is only accepted when sent as a COMMAND_LONG.
	MAV_RESULT_COMMAND_LONG_ONLY MAV_RESULT = common.MAV_RESULT_COMMAND_LONG_ONLY
	// Command is only accepted when sent as a COMMAND_INT.
	MAV_RESULT_COMMAND_INT_ONLY MAV_RESULT = common.MAV_RESULT_COMMAND_INT_ONLY
	// Command is invalid because a frame is required and the specified frame is not supported.
	MAV_RESULT_COMMAND_UNSUPPORTED_MAV_FRAME MAV_RESULT = common.MAV_RESULT_COMMAND_UNSUPPORTED_MAV_FRAME
	// Sender is not authorized to control this MAV component. Control may be requested using MAV_CMD_REQUEST_OPERATOR_CONTROL.
	MAV_RESULT_PERMISSION_DENIED MAV_RESULT = common.MAV_RESULT_PERMISSION_DENIED
)
