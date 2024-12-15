//autogenerated:yes
//nolint:revive,misspell,govet,lll
package common

// Fuel status.
// This message provides "generic" fuel level information for display in a GCS and for triggering failsafes in an autopilot.
// The fuel type and associated units for fields in this message are defined in the enum MAV_FUEL_TYPE.
// The reported `consumed_fuel` and `remaining_fuel` must only be supplied if measured: they must not be inferred from the `maximum_fuel` and the other value.
// A recipient can assume that if these fields are supplied they are accurate.
// If not provided, the recipient can infer `remaining_fuel` from `maximum_fuel` and `consumed_fuel` on the assumption that the fuel was initially at its maximum (this is what battery monitors assume).
// Note however that this is an assumption, and the UI should prompt the user appropriately (i.e. notify user that they should fill the tank before boot).
// This kind of information may also be sent in fuel-specific messages such as BATTERY_STATUS_V2.
// If both messages are sent for the same fuel system, the ids and corresponding information must match.
// This should be streamed (nominally at 0.1 Hz).
type MessageFuelStatus struct {
	// Fuel ID. Must match ID of other messages for same fuel system, such as BATTERY_STATUS_V2.
	Id uint8
	// Capacity when full. Must be provided.
	MaximumFuel float32
	// Consumed fuel (measured). This value should not be inferred: if not measured set to NaN. NaN: field not provided.
	ConsumedFuel float32
	// Remaining fuel until empty (measured). The value should not be inferred: if not measured set to NaN. NaN: field not provided.
	RemainingFuel float32
	// Percentage of remaining fuel, relative to full. Values: [0-100], UINT8_MAX: field not provided.
	PercentRemaining uint8
	// Positive value when emptying/using, and negative if filling/replacing. NaN: field not provided.
	FlowRate float32
	// Fuel temperature. NaN: field not provided.
	Temperature float32
	// Fuel type. Defines units for fuel capacity and consumption fields above.
	FuelType MAV_FUEL_TYPE `mavenum:"uint32"`
}

// GetID implements the message.Message interface.
func (*MessageFuelStatus) GetID() uint32 {
	return 371
}