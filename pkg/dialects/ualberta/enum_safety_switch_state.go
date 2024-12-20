//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Possible safety switch states.
type SAFETY_SWITCH_STATE = common.SAFETY_SWITCH_STATE

const (
	// Safety switch is engaged and vehicle should be safe to approach.
	SAFETY_SWITCH_STATE_SAFE SAFETY_SWITCH_STATE = common.SAFETY_SWITCH_STATE_SAFE
	// Safety switch is NOT engaged and motors, propellers and other actuators should be considered active.
	SAFETY_SWITCH_STATE_DANGEROUS SAFETY_SWITCH_STATE = common.SAFETY_SWITCH_STATE_DANGEROUS
)
