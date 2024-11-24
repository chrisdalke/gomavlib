//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// These flags are used to diagnose the failure state of CELLULAR_STATUS
type CELLULAR_NETWORK_FAILED_REASON = common.CELLULAR_NETWORK_FAILED_REASON

const (
	// No error
	CELLULAR_NETWORK_FAILED_REASON_NONE CELLULAR_NETWORK_FAILED_REASON = common.CELLULAR_NETWORK_FAILED_REASON_NONE
	// Error state is unknown
	CELLULAR_NETWORK_FAILED_REASON_UNKNOWN CELLULAR_NETWORK_FAILED_REASON = common.CELLULAR_NETWORK_FAILED_REASON_UNKNOWN
	// SIM is required for the modem but missing
	CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING CELLULAR_NETWORK_FAILED_REASON = common.CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING
	// SIM is available, but not usable for connection
	CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR CELLULAR_NETWORK_FAILED_REASON = common.CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR
)
