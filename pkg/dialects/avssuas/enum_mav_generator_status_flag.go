//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Flags to report status/failure cases for a power generator (used in GENERATOR_STATUS). Note that FAULTS are conditions that cause the generator to fail. Warnings are conditions that require attention before the next use (they indicate the system is not operating properly).
type MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG

const (
	// Generator is off.
	MAV_GENERATOR_STATUS_FLAG_OFF MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_OFF
	// Generator is ready to start generating power.
	MAV_GENERATOR_STATUS_FLAG_READY MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_READY
	// Generator is generating power.
	MAV_GENERATOR_STATUS_FLAG_GENERATING MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_GENERATING
	// Generator is charging the batteries (generating enough power to charge and provide the load).
	MAV_GENERATOR_STATUS_FLAG_CHARGING MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_CHARGING
	// Generator is operating at a reduced maximum power.
	MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER
	// Generator is providing the maximum output.
	MAV_GENERATOR_STATUS_FLAG_MAXPOWER MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_MAXPOWER
	// Generator is near the maximum operating temperature, cooling is insufficient.
	MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING
	// Generator hit the maximum operating temperature and shutdown.
	MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT
	// Power electronics are near the maximum operating temperature, cooling is insufficient.
	MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING
	// Power electronics hit the maximum operating temperature and shutdown.
	MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT
	// Power electronics experienced a fault and shutdown.
	MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT
	// The power source supplying the generator failed e.g. mechanical generator stopped, tether is no longer providing power, solar cell is in shade, hydrogen reaction no longer happening.
	MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT
	// Generator controller having communication problems.
	MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING
	// Power electronic or generator cooling system error.
	MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING
	// Generator controller power rail experienced a fault.
	MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT
	// Generator controller exceeded the overcurrent threshold and shutdown to prevent damage.
	MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT
	// Generator controller detected a high current going into the batteries and shutdown to prevent battery damage.
	MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT
	// Generator controller exceeded it's overvoltage threshold and shutdown to prevent it exceeding the voltage rating.
	MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT
	// Batteries are under voltage (generator will not start).
	MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT
	// Generator start is inhibited by e.g. a safety switch.
	MAV_GENERATOR_STATUS_FLAG_START_INHIBITED MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_START_INHIBITED
	// Generator requires maintenance.
	MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED
	// Generator is not ready to generate yet.
	MAV_GENERATOR_STATUS_FLAG_WARMING_UP MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_WARMING_UP
	// Generator is idle.
	MAV_GENERATOR_STATUS_FLAG_IDLE MAV_GENERATOR_STATUS_FLAG = common.MAV_GENERATOR_STATUS_FLAG_IDLE
)
