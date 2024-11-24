//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// These flags encode the cellular network status
type CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG

const (
	// State unknown or not reportable.
	CELLULAR_STATUS_FLAG_UNKNOWN CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_UNKNOWN
	// Modem is unusable
	CELLULAR_STATUS_FLAG_FAILED CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_FAILED
	// Modem is being initialized
	CELLULAR_STATUS_FLAG_INITIALIZING CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_INITIALIZING
	// Modem is locked
	CELLULAR_STATUS_FLAG_LOCKED CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_LOCKED
	// Modem is not enabled and is powered down
	CELLULAR_STATUS_FLAG_DISABLED CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_DISABLED
	// Modem is currently transitioning to the CELLULAR_STATUS_FLAG_DISABLED state
	CELLULAR_STATUS_FLAG_DISABLING CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_DISABLING
	// Modem is currently transitioning to the CELLULAR_STATUS_FLAG_ENABLED state
	CELLULAR_STATUS_FLAG_ENABLING CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_ENABLING
	// Modem is enabled and powered on but not registered with a network provider and not available for data connections
	CELLULAR_STATUS_FLAG_ENABLED CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_ENABLED
	// Modem is searching for a network provider to register
	CELLULAR_STATUS_FLAG_SEARCHING CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_SEARCHING
	// Modem is registered with a network provider, and data connections and messaging may be available for use
	CELLULAR_STATUS_FLAG_REGISTERED CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_REGISTERED
	// Modem is disconnecting and deactivating the last active packet data bearer. This state will not be entered if more than one packet data bearer is active and one of the active bearers is deactivated
	CELLULAR_STATUS_FLAG_DISCONNECTING CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_DISCONNECTING
	// Modem is activating and connecting the first packet data bearer. Subsequent bearer activations when another bearer is already active do not cause this state to be entered
	CELLULAR_STATUS_FLAG_CONNECTING CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_CONNECTING
	// One or more packet data bearers is active and connected
	CELLULAR_STATUS_FLAG_CONNECTED CELLULAR_STATUS_FLAG = common.CELLULAR_STATUS_FLAG_CONNECTED
)
