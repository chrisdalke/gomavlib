//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// List of possible units where failures can be injected.
type FAILURE_UNIT = common.FAILURE_UNIT

const (
	FAILURE_UNIT_SENSOR_GYRO            FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_GYRO
	FAILURE_UNIT_SENSOR_ACCEL           FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_ACCEL
	FAILURE_UNIT_SENSOR_MAG             FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_MAG
	FAILURE_UNIT_SENSOR_BARO            FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_BARO
	FAILURE_UNIT_SENSOR_GPS             FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_GPS
	FAILURE_UNIT_SENSOR_OPTICAL_FLOW    FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_OPTICAL_FLOW
	FAILURE_UNIT_SENSOR_VIO             FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_VIO
	FAILURE_UNIT_SENSOR_DISTANCE_SENSOR FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_DISTANCE_SENSOR
	FAILURE_UNIT_SENSOR_AIRSPEED        FAILURE_UNIT = common.FAILURE_UNIT_SENSOR_AIRSPEED
	FAILURE_UNIT_SYSTEM_BATTERY         FAILURE_UNIT = common.FAILURE_UNIT_SYSTEM_BATTERY
	FAILURE_UNIT_SYSTEM_MOTOR           FAILURE_UNIT = common.FAILURE_UNIT_SYSTEM_MOTOR
	FAILURE_UNIT_SYSTEM_SERVO           FAILURE_UNIT = common.FAILURE_UNIT_SYSTEM_SERVO
	FAILURE_UNIT_SYSTEM_AVOIDANCE       FAILURE_UNIT = common.FAILURE_UNIT_SYSTEM_AVOIDANCE
	FAILURE_UNIT_SYSTEM_RC_SIGNAL       FAILURE_UNIT = common.FAILURE_UNIT_SYSTEM_RC_SIGNAL
	FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL  FAILURE_UNIT = common.FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL
)
