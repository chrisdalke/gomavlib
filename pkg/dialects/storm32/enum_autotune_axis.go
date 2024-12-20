//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Enable axes that will be tuned via autotuning. Used in MAV_CMD_DO_AUTOTUNE_ENABLE.
type AUTOTUNE_AXIS = common.AUTOTUNE_AXIS

const (
	// Flight stack tunes axis according to its default settings.
	AUTOTUNE_AXIS_DEFAULT AUTOTUNE_AXIS = common.AUTOTUNE_AXIS_DEFAULT
	// Autotune roll axis.
	AUTOTUNE_AXIS_ROLL AUTOTUNE_AXIS = common.AUTOTUNE_AXIS_ROLL
	// Autotune pitch axis.
	AUTOTUNE_AXIS_PITCH AUTOTUNE_AXIS = common.AUTOTUNE_AXIS_PITCH
	// Autotune yaw axis.
	AUTOTUNE_AXIS_YAW AUTOTUNE_AXIS = common.AUTOTUNE_AXIS_YAW
)
