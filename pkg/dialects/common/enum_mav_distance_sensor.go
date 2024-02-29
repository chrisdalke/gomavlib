//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Enumeration of distance sensor types
type MAV_DISTANCE_SENSOR uint64

const (
	// Laser rangefinder, e.g. LightWare SF02/F or PulsedLight units
	MAV_DISTANCE_SENSOR_LASER MAV_DISTANCE_SENSOR = 0
	// Ultrasound rangefinder, e.g. MaxBotix units
	MAV_DISTANCE_SENSOR_ULTRASOUND MAV_DISTANCE_SENSOR = 1
	// Infrared rangefinder, e.g. Sharp units
	MAV_DISTANCE_SENSOR_INFRARED MAV_DISTANCE_SENSOR = 2
	// Radar type, e.g. uLanding units
	MAV_DISTANCE_SENSOR_RADAR MAV_DISTANCE_SENSOR = 3
	// Broken or unknown type, e.g. analog units
	MAV_DISTANCE_SENSOR_UNKNOWN MAV_DISTANCE_SENSOR = 4
)

var labels_MAV_DISTANCE_SENSOR = map[MAV_DISTANCE_SENSOR]string{
	MAV_DISTANCE_SENSOR_LASER:      "MAV_DISTANCE_SENSOR_LASER",
	MAV_DISTANCE_SENSOR_ULTRASOUND: "MAV_DISTANCE_SENSOR_ULTRASOUND",
	MAV_DISTANCE_SENSOR_INFRARED:   "MAV_DISTANCE_SENSOR_INFRARED",
	MAV_DISTANCE_SENSOR_RADAR:      "MAV_DISTANCE_SENSOR_RADAR",
	MAV_DISTANCE_SENSOR_UNKNOWN:    "MAV_DISTANCE_SENSOR_UNKNOWN",
}

var values_MAV_DISTANCE_SENSOR = map[string]MAV_DISTANCE_SENSOR{
	"MAV_DISTANCE_SENSOR_LASER":      MAV_DISTANCE_SENSOR_LASER,
	"MAV_DISTANCE_SENSOR_ULTRASOUND": MAV_DISTANCE_SENSOR_ULTRASOUND,
	"MAV_DISTANCE_SENSOR_INFRARED":   MAV_DISTANCE_SENSOR_INFRARED,
	"MAV_DISTANCE_SENSOR_RADAR":      MAV_DISTANCE_SENSOR_RADAR,
	"MAV_DISTANCE_SENSOR_UNKNOWN":    MAV_DISTANCE_SENSOR_UNKNOWN,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_DISTANCE_SENSOR) MarshalText() ([]byte, error) {
	if name, ok := labels_MAV_DISTANCE_SENSOR[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_DISTANCE_SENSOR) UnmarshalText(text []byte) error {
	if value, ok := values_MAV_DISTANCE_SENSOR[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = MAV_DISTANCE_SENSOR(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_DISTANCE_SENSOR) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
