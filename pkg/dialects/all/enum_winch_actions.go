//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Winch actions.
type WINCH_ACTIONS = common.WINCH_ACTIONS

const (
	// Allow motor to freewheel.
	WINCH_RELAXED WINCH_ACTIONS = common.WINCH_RELAXED
	// Wind or unwind specified length of line, optionally using specified rate.
	WINCH_RELATIVE_LENGTH_CONTROL WINCH_ACTIONS = common.WINCH_RELATIVE_LENGTH_CONTROL
	// Wind or unwind line at specified rate.
	WINCH_RATE_CONTROL WINCH_ACTIONS = common.WINCH_RATE_CONTROL
	// Perform the locking sequence to relieve motor while in the fully retracted position. Only action and instance command parameters are used, others are ignored.
	WINCH_LOCK WINCH_ACTIONS = common.WINCH_LOCK
	// Sequence of drop, slow down, touch down, reel up, lock. Only action and instance command parameters are used, others are ignored.
	WINCH_DELIVER WINCH_ACTIONS = common.WINCH_DELIVER
	// Engage motor and hold current position. Only action and instance command parameters are used, others are ignored.
	WINCH_HOLD WINCH_ACTIONS = common.WINCH_HOLD
	// Return the reel to the fully retracted position. Only action and instance command parameters are used, others are ignored.
	WINCH_RETRACT WINCH_ACTIONS = common.WINCH_RETRACT
	// Load the reel with line. The winch will calculate the total loaded length and stop when the tension exceeds a threshold. Only action and instance command parameters are used, others are ignored.
	WINCH_LOAD_LINE WINCH_ACTIONS = common.WINCH_LOAD_LINE
	// Spool out the entire length of the line. Only action and instance command parameters are used, others are ignored.
	WINCH_ABANDON_LINE WINCH_ACTIONS = common.WINCH_ABANDON_LINE
	// Spools out just enough to present the hook to the user to load the payload. Only action and instance command parameters are used, others are ignored
	WINCH_LOAD_PAYLOAD WINCH_ACTIONS = common.WINCH_LOAD_PAYLOAD
)
