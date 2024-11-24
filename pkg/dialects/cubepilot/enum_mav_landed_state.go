//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Enumeration of landed detector states
type MAV_LANDED_STATE = common.MAV_LANDED_STATE

const (
	// MAV landed state is unknown
	MAV_LANDED_STATE_UNDEFINED MAV_LANDED_STATE = common.MAV_LANDED_STATE_UNDEFINED
	// MAV is landed (on ground)
	MAV_LANDED_STATE_ON_GROUND MAV_LANDED_STATE = common.MAV_LANDED_STATE_ON_GROUND
	// MAV is in air
	MAV_LANDED_STATE_IN_AIR MAV_LANDED_STATE = common.MAV_LANDED_STATE_IN_AIR
	// MAV currently taking off
	MAV_LANDED_STATE_TAKEOFF MAV_LANDED_STATE = common.MAV_LANDED_STATE_TAKEOFF
	// MAV currently landing
	MAV_LANDED_STATE_LANDING MAV_LANDED_STATE = common.MAV_LANDED_STATE_LANDING
)
