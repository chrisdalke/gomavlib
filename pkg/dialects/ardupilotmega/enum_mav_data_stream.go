//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// A data stream is not a fixed set of messages, but rather a
// recommendation to the autopilot software. Individual autopilots may or may not obey
// the recommended messages.
type MAV_DATA_STREAM = common.MAV_DATA_STREAM

const (
	// Enable all data streams
	MAV_DATA_STREAM_ALL MAV_DATA_STREAM = common.MAV_DATA_STREAM_ALL
	// Enable IMU_RAW, GPS_RAW, GPS_STATUS packets.
	MAV_DATA_STREAM_RAW_SENSORS MAV_DATA_STREAM = common.MAV_DATA_STREAM_RAW_SENSORS
	// Enable GPS_STATUS, CONTROL_STATUS, AUX_STATUS
	MAV_DATA_STREAM_EXTENDED_STATUS MAV_DATA_STREAM = common.MAV_DATA_STREAM_EXTENDED_STATUS
	// Enable RC_CHANNELS_SCALED, RC_CHANNELS_RAW, SERVO_OUTPUT_RAW
	MAV_DATA_STREAM_RC_CHANNELS MAV_DATA_STREAM = common.MAV_DATA_STREAM_RC_CHANNELS
	// Enable ATTITUDE_CONTROLLER_OUTPUT, POSITION_CONTROLLER_OUTPUT, NAV_CONTROLLER_OUTPUT.
	MAV_DATA_STREAM_RAW_CONTROLLER MAV_DATA_STREAM = common.MAV_DATA_STREAM_RAW_CONTROLLER
	// Enable LOCAL_POSITION, GLOBAL_POSITION_INT messages.
	MAV_DATA_STREAM_POSITION MAV_DATA_STREAM = common.MAV_DATA_STREAM_POSITION
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA1 MAV_DATA_STREAM = common.MAV_DATA_STREAM_EXTRA1
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA2 MAV_DATA_STREAM = common.MAV_DATA_STREAM_EXTRA2
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA3 MAV_DATA_STREAM = common.MAV_DATA_STREAM_EXTRA3
)
