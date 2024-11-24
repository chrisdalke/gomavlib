//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Speed setpoint types used in MAV_CMD_DO_CHANGE_SPEED
type SPEED_TYPE = common.SPEED_TYPE

const (
	// Airspeed
	SPEED_TYPE_AIRSPEED SPEED_TYPE = common.SPEED_TYPE_AIRSPEED
	// Groundspeed
	SPEED_TYPE_GROUNDSPEED SPEED_TYPE = common.SPEED_TYPE_GROUNDSPEED
	// Climb speed
	SPEED_TYPE_CLIMB_SPEED SPEED_TYPE = common.SPEED_TYPE_CLIMB_SPEED
	// Descent speed
	SPEED_TYPE_DESCENT_SPEED SPEED_TYPE = common.SPEED_TYPE_DESCENT_SPEED
)
