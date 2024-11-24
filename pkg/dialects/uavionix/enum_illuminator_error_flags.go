//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Illuminator module error flags (bitmap, 0 means no error)
type ILLUMINATOR_ERROR_FLAGS = common.ILLUMINATOR_ERROR_FLAGS

const (
	// Illuminator thermal throttling error.
	ILLUMINATOR_ERROR_FLAGS_THERMAL_THROTTLING ILLUMINATOR_ERROR_FLAGS = common.ILLUMINATOR_ERROR_FLAGS_THERMAL_THROTTLING
	// Illuminator over temperature shutdown error.
	ILLUMINATOR_ERROR_FLAGS_OVER_TEMPERATURE_SHUTDOWN ILLUMINATOR_ERROR_FLAGS = common.ILLUMINATOR_ERROR_FLAGS_OVER_TEMPERATURE_SHUTDOWN
	// Illuminator thermistor failure.
	ILLUMINATOR_ERROR_FLAGS_THERMISTOR_FAILURE ILLUMINATOR_ERROR_FLAGS = common.ILLUMINATOR_ERROR_FLAGS_THERMISTOR_FAILURE
)
