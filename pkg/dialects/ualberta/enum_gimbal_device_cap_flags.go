//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Gimbal device (low level) capability flags (bitmap).
type GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS

const (
	// Gimbal device supports a retracted position.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT
	// Gimbal device supports a horizontal, forward looking position, stabilized.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL
	// Gimbal device supports rotating around roll axis.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS
	// Gimbal device supports to follow a roll angle relative to the vehicle.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW
	// Gimbal device supports locking to a roll angle (generally that's the default with roll stabilized).
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK
	// Gimbal device supports rotating around pitch axis.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS
	// Gimbal device supports to follow a pitch angle relative to the vehicle.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW
	// Gimbal device supports locking to a pitch angle (generally that's the default with pitch stabilized).
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK
	// Gimbal device supports rotating around yaw axis.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS
	// Gimbal device supports to follow a yaw angle relative to the vehicle (generally that's the default).
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW
	// Gimbal device supports locking to an absolute heading, i.e., yaw angle relative to North (earth frame, often this is an option available).
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK
	// Gimbal device supports yawing/panning infinitely (e.g. using slip disk).
	GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW
	// Gimbal device supports yaw angles and angular velocities relative to North (earth frame). This usually requires support by an autopilot via AUTOPILOT_STATE_FOR_GIMBAL_DEVICE. Support can go on and off during runtime, which is reported by the flag GIMBAL_DEVICE_FLAGS_CAN_ACCEPT_YAW_IN_EARTH_FRAME.
	GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME
	// Gimbal device supports radio control inputs as an alternative input for controlling the gimbal orientation.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_RC_INPUTS GIMBAL_DEVICE_CAP_FLAGS = common.GIMBAL_DEVICE_CAP_FLAGS_HAS_RC_INPUTS
)
