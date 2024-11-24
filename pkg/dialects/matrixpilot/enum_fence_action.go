//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package matrixpilot

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Actions following geofence breach.
type FENCE_ACTION = common.FENCE_ACTION

const (
	// Disable fenced mode. If used in a plan this would mean the next fence is disabled.
	FENCE_ACTION_NONE FENCE_ACTION = common.FENCE_ACTION_NONE
	// Fly to geofence MAV_CMD_NAV_FENCE_RETURN_POINT in GUIDED mode. Note: This action is only supported by ArduPlane, and may not be supported in all versions.
	FENCE_ACTION_GUIDED FENCE_ACTION = common.FENCE_ACTION_GUIDED
	// Report fence breach, but don't take action
	FENCE_ACTION_REPORT FENCE_ACTION = common.FENCE_ACTION_REPORT
	// Fly to geofence MAV_CMD_NAV_FENCE_RETURN_POINT with manual throttle control in GUIDED mode. Note: This action is only supported by ArduPlane, and may not be supported in all versions.
	FENCE_ACTION_GUIDED_THR_PASS FENCE_ACTION = common.FENCE_ACTION_GUIDED_THR_PASS
	// Return/RTL mode.
	FENCE_ACTION_RTL FENCE_ACTION = common.FENCE_ACTION_RTL
	// Hold at current location.
	FENCE_ACTION_HOLD FENCE_ACTION = common.FENCE_ACTION_HOLD
	// Termination failsafe. Motors are shut down (some flight stacks may trigger other failsafe actions).
	FENCE_ACTION_TERMINATE FENCE_ACTION = common.FENCE_ACTION_TERMINATE
	// Land at current location.
	FENCE_ACTION_LAND FENCE_ACTION = common.FENCE_ACTION_LAND
)
