//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Enumeration of VTOL states
type MAV_VTOL_STATE = common.MAV_VTOL_STATE

const (
	// MAV is not configured as VTOL
	MAV_VTOL_STATE_UNDEFINED MAV_VTOL_STATE = common.MAV_VTOL_STATE_UNDEFINED
	// VTOL is in transition from multicopter to fixed-wing
	MAV_VTOL_STATE_TRANSITION_TO_FW MAV_VTOL_STATE = common.MAV_VTOL_STATE_TRANSITION_TO_FW
	// VTOL is in transition from fixed-wing to multicopter
	MAV_VTOL_STATE_TRANSITION_TO_MC MAV_VTOL_STATE = common.MAV_VTOL_STATE_TRANSITION_TO_MC
	// VTOL is in multicopter state
	MAV_VTOL_STATE_MC MAV_VTOL_STATE = common.MAV_VTOL_STATE_MC
	// VTOL is in fixed-wing state
	MAV_VTOL_STATE_FW MAV_VTOL_STATE = common.MAV_VTOL_STATE_FW
)
