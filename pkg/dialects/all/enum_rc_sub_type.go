//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// RC sub-type of types defined in RC_TYPE. Used in MAV_CMD_START_RX_PAIR. Ignored if value does not correspond to the set RC_TYPE.
type RC_SUB_TYPE = common.RC_SUB_TYPE

const (
	// Spektrum DSM2
	RC_SUB_TYPE_SPEKTRUM_DSM2 RC_SUB_TYPE = common.RC_SUB_TYPE_SPEKTRUM_DSM2
	// Spektrum DSMX
	RC_SUB_TYPE_SPEKTRUM_DSMX RC_SUB_TYPE = common.RC_SUB_TYPE_SPEKTRUM_DSMX
	// Spektrum DSMX8
	RC_SUB_TYPE_SPEKTRUM_DSMX8 RC_SUB_TYPE = common.RC_SUB_TYPE_SPEKTRUM_DSMX8
)
