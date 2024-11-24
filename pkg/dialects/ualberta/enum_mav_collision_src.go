//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Source of information about this collision.
type MAV_COLLISION_SRC = common.MAV_COLLISION_SRC

const (
	// ID field references ADSB_VEHICLE packets
	MAV_COLLISION_SRC_ADSB MAV_COLLISION_SRC = common.MAV_COLLISION_SRC_ADSB
	// ID field references MAVLink SRC ID
	MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT MAV_COLLISION_SRC = common.MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT
)
