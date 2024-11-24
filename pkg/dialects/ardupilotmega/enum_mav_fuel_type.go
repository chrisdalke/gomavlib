//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Fuel types for use in FUEL_TYPE. Fuel types specify the units for the maximum, available and consumed fuel, and for the flow rates.
type MAV_FUEL_TYPE = common.MAV_FUEL_TYPE

const (
	// Not specified. Fuel levels are normalized (i.e. maximum is 1, and other levels are relative to 1).
	MAV_FUEL_TYPE_UNKNOWN MAV_FUEL_TYPE = common.MAV_FUEL_TYPE_UNKNOWN
	// A generic liquid fuel. Fuel levels are in millilitres (ml). Fuel rates are in millilitres/second.
	MAV_FUEL_TYPE_LIQUID MAV_FUEL_TYPE = common.MAV_FUEL_TYPE_LIQUID
	// A gas tank. Fuel levels are in kilo-Pascal (kPa), and flow rates are in milliliters per second (ml/s).
	MAV_FUEL_TYPE_GAS MAV_FUEL_TYPE = common.MAV_FUEL_TYPE_GAS
)
