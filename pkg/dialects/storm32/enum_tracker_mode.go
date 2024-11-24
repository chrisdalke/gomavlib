//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/ardupilotmega"
)

// A mapping of antenna tracker flight modes for custom_mode field of heartbeat.
type TRACKER_MODE = ardupilotmega.TRACKER_MODE

const (
	TRACKER_MODE_MANUAL       TRACKER_MODE = ardupilotmega.TRACKER_MODE_MANUAL
	TRACKER_MODE_STOP         TRACKER_MODE = ardupilotmega.TRACKER_MODE_STOP
	TRACKER_MODE_SCAN         TRACKER_MODE = ardupilotmega.TRACKER_MODE_SCAN
	TRACKER_MODE_SERVO_TEST   TRACKER_MODE = ardupilotmega.TRACKER_MODE_SERVO_TEST
	TRACKER_MODE_AUTO         TRACKER_MODE = ardupilotmega.TRACKER_MODE_AUTO
	TRACKER_MODE_INITIALIZING TRACKER_MODE = ardupilotmega.TRACKER_MODE_INITIALIZING
)
