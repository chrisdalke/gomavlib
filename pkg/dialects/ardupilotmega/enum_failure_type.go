//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// List of possible failure type to inject.
type FAILURE_TYPE = common.FAILURE_TYPE

const (
	// No failure injected, used to reset a previous failure.
	FAILURE_TYPE_OK FAILURE_TYPE = common.FAILURE_TYPE_OK
	// Sets unit off, so completely non-responsive.
	FAILURE_TYPE_OFF FAILURE_TYPE = common.FAILURE_TYPE_OFF
	// Unit is stuck e.g. keeps reporting the same value.
	FAILURE_TYPE_STUCK FAILURE_TYPE = common.FAILURE_TYPE_STUCK
	// Unit is reporting complete garbage.
	FAILURE_TYPE_GARBAGE FAILURE_TYPE = common.FAILURE_TYPE_GARBAGE
	// Unit is consistently wrong.
	FAILURE_TYPE_WRONG FAILURE_TYPE = common.FAILURE_TYPE_WRONG
	// Unit is slow, so e.g. reporting at slower than expected rate.
	FAILURE_TYPE_SLOW FAILURE_TYPE = common.FAILURE_TYPE_SLOW
	// Data of unit is delayed in time.
	FAILURE_TYPE_DELAYED FAILURE_TYPE = common.FAILURE_TYPE_DELAYED
	// Unit is sometimes working, sometimes not.
	FAILURE_TYPE_INTERMITTENT FAILURE_TYPE = common.FAILURE_TYPE_INTERMITTENT
)
