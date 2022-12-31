//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"errors"
)

// Smart battery supply status/fault flags (bitmask) for health indication. The battery must also report either MAV_BATTERY_CHARGE_STATE_FAILED or MAV_BATTERY_CHARGE_STATE_UNHEALTHY if any of these are set.
type MAV_BATTERY_FAULT uint32

const (
	// Battery has deep discharged.
	MAV_BATTERY_FAULT_DEEP_DISCHARGE MAV_BATTERY_FAULT = 1
	// Voltage spikes.
	MAV_BATTERY_FAULT_SPIKES MAV_BATTERY_FAULT = 2
	// One or more cells have failed. Battery should also report MAV_BATTERY_CHARGE_STATE_FAILE (and should not be used).
	MAV_BATTERY_FAULT_CELL_FAIL MAV_BATTERY_FAULT = 4
	// Over-current fault.
	MAV_BATTERY_FAULT_OVER_CURRENT MAV_BATTERY_FAULT = 8
	// Over-temperature fault.
	MAV_BATTERY_FAULT_OVER_TEMPERATURE MAV_BATTERY_FAULT = 16
	// Under-temperature fault.
	MAV_BATTERY_FAULT_UNDER_TEMPERATURE MAV_BATTERY_FAULT = 32
	// Vehicle voltage is not compatible with this battery (batteries on same power rail should have similar voltage).
	MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE MAV_BATTERY_FAULT = 64
	// Battery firmware is not compatible with current autopilot firmware.
	MAV_BATTERY_FAULT_INCOMPATIBLE_FIRMWARE MAV_BATTERY_FAULT = 128
	// Battery is not compatible due to cell configuration (e.g. 5s1p when vehicle requires 6s).
	BATTERY_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION MAV_BATTERY_FAULT = 256
)

var labels_MAV_BATTERY_FAULT = map[MAV_BATTERY_FAULT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_FAULT) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_BATTERY_FAULT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_BATTERY_FAULT = map[string]MAV_BATTERY_FAULT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_FAULT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_BATTERY_FAULT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_BATTERY_FAULT) String() string {
	if l, ok := labels_MAV_BATTERY_FAULT[e]; ok {
		return l
	}
	return "invalid value"
}