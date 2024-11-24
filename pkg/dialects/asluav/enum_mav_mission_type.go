//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Type of mission items being requested/sent in mission protocol.
type MAV_MISSION_TYPE = common.MAV_MISSION_TYPE

const (
	// Items are mission commands for main mission.
	MAV_MISSION_TYPE_MISSION MAV_MISSION_TYPE = common.MAV_MISSION_TYPE_MISSION
	// Specifies GeoFence area(s). Items are MAV_CMD_NAV_FENCE_ GeoFence items.
	MAV_MISSION_TYPE_FENCE MAV_MISSION_TYPE = common.MAV_MISSION_TYPE_FENCE
	// Specifies the rally points for the vehicle. Rally points are alternative RTL points. Items are MAV_CMD_NAV_RALLY_POINT rally point items.
	MAV_MISSION_TYPE_RALLY MAV_MISSION_TYPE = common.MAV_MISSION_TYPE_RALLY
	// Only used in MISSION_CLEAR_ALL to clear all mission types.
	MAV_MISSION_TYPE_ALL MAV_MISSION_TYPE = common.MAV_MISSION_TYPE_ALL
)
