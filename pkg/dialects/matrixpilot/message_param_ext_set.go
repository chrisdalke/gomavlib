//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Set a parameter value. In order to deal with message loss (and retransmission of PARAM_EXT_SET), when setting a parameter value and the new value is the same as the current value, you will immediately get a PARAM_ACK_ACCEPTED response. If the current state is PARAM_ACK_IN_PROGRESS, you will accordingly receive a PARAM_ACK_IN_PROGRESS in response.
type MessageParamExtSet = common.MessageParamExtSet
