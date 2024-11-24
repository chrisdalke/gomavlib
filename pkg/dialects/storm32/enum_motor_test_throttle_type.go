//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Defines how throttle value is represented in MAV_CMD_DO_MOTOR_TEST.
type MOTOR_TEST_THROTTLE_TYPE = common.MOTOR_TEST_THROTTLE_TYPE

const (
	// Throttle as a percentage (0 ~ 100)
	MOTOR_TEST_THROTTLE_PERCENT MOTOR_TEST_THROTTLE_TYPE = common.MOTOR_TEST_THROTTLE_PERCENT
	// Throttle as an absolute PWM value (normally in range of 1000~2000).
	MOTOR_TEST_THROTTLE_PWM MOTOR_TEST_THROTTLE_TYPE = common.MOTOR_TEST_THROTTLE_PWM
	// Throttle pass-through from pilot's transmitter.
	MOTOR_TEST_THROTTLE_PILOT MOTOR_TEST_THROTTLE_TYPE = common.MOTOR_TEST_THROTTLE_PILOT
	// Per-motor compass calibration test.
	MOTOR_TEST_COMPASS_CAL MOTOR_TEST_THROTTLE_TYPE = common.MOTOR_TEST_COMPASS_CAL
)
