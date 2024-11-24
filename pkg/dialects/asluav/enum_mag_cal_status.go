//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

type MAG_CAL_STATUS = common.MAG_CAL_STATUS

const (
	MAG_CAL_NOT_STARTED      MAG_CAL_STATUS = common.MAG_CAL_NOT_STARTED
	MAG_CAL_WAITING_TO_START MAG_CAL_STATUS = common.MAG_CAL_WAITING_TO_START
	MAG_CAL_RUNNING_STEP_ONE MAG_CAL_STATUS = common.MAG_CAL_RUNNING_STEP_ONE
	MAG_CAL_RUNNING_STEP_TWO MAG_CAL_STATUS = common.MAG_CAL_RUNNING_STEP_TWO
	MAG_CAL_SUCCESS          MAG_CAL_STATUS = common.MAG_CAL_SUCCESS
	MAG_CAL_FAILED           MAG_CAL_STATUS = common.MAG_CAL_FAILED
	MAG_CAL_BAD_ORIENTATION  MAG_CAL_STATUS = common.MAG_CAL_BAD_ORIENTATION
	MAG_CAL_BAD_RADIUS       MAG_CAL_STATUS = common.MAG_CAL_BAD_RADIUS
)
