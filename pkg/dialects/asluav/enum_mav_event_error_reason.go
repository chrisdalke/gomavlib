//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Reason for an event error response.
type MAV_EVENT_ERROR_REASON = common.MAV_EVENT_ERROR_REASON

const (
	// The requested event is not available (anymore).
	MAV_EVENT_ERROR_REASON_UNAVAILABLE MAV_EVENT_ERROR_REASON = common.MAV_EVENT_ERROR_REASON_UNAVAILABLE
)
