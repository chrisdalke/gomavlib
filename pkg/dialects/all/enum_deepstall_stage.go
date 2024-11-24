//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/ardupilotmega"
)

// Deepstall flight stage.
type DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE

const (
	// Flying to the landing point.
	DEEPSTALL_STAGE_FLY_TO_LANDING DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE_FLY_TO_LANDING
	// Building an estimate of the wind.
	DEEPSTALL_STAGE_ESTIMATE_WIND DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE_ESTIMATE_WIND
	// Waiting to breakout of the loiter to fly the approach.
	DEEPSTALL_STAGE_WAIT_FOR_BREAKOUT DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE_WAIT_FOR_BREAKOUT
	// Flying to the first arc point to turn around to the landing point.
	DEEPSTALL_STAGE_FLY_TO_ARC DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE_FLY_TO_ARC
	// Turning around back to the deepstall landing point.
	DEEPSTALL_STAGE_ARC DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE_ARC
	// Approaching the landing point.
	DEEPSTALL_STAGE_APPROACH DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE_APPROACH
	// Stalling and steering towards the land point.
	DEEPSTALL_STAGE_LAND DEEPSTALL_STAGE = ardupilotmega.DEEPSTALL_STAGE_LAND
)
