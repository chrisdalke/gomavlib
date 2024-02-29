//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
	"strings"
)

// These flags are used in the AIS_VESSEL.fields bitmask to indicate validity of data in the other message fields. When set, the data is valid.
type AIS_FLAGS uint64

const (
	// 1 = Position accuracy less than 10m, 0 = position accuracy greater than 10m.
	AIS_FLAGS_POSITION_ACCURACY AIS_FLAGS = 1
	AIS_FLAGS_VALID_COG         AIS_FLAGS = 2
	AIS_FLAGS_VALID_VELOCITY    AIS_FLAGS = 4
	// 1 = Velocity over 52.5765m/s (102.2 knots)
	AIS_FLAGS_HIGH_VELOCITY   AIS_FLAGS = 8
	AIS_FLAGS_VALID_TURN_RATE AIS_FLAGS = 16
	// Only the sign of the returned turn rate value is valid, either greater than 5deg/30s or less than -5deg/30s
	AIS_FLAGS_TURN_RATE_SIGN_ONLY AIS_FLAGS = 32
	AIS_FLAGS_VALID_DIMENSIONS    AIS_FLAGS = 64
	// Distance to bow is larger than 511m
	AIS_FLAGS_LARGE_BOW_DIMENSION AIS_FLAGS = 128
	// Distance to stern is larger than 511m
	AIS_FLAGS_LARGE_STERN_DIMENSION AIS_FLAGS = 256
	// Distance to port side is larger than 63m
	AIS_FLAGS_LARGE_PORT_DIMENSION AIS_FLAGS = 512
	// Distance to starboard side is larger than 63m
	AIS_FLAGS_LARGE_STARBOARD_DIMENSION AIS_FLAGS = 1024
	AIS_FLAGS_VALID_CALLSIGN            AIS_FLAGS = 2048
	AIS_FLAGS_VALID_NAME                AIS_FLAGS = 4096
)

var labels_AIS_FLAGS = map[AIS_FLAGS]string{
	AIS_FLAGS_POSITION_ACCURACY:         "AIS_FLAGS_POSITION_ACCURACY",
	AIS_FLAGS_VALID_COG:                 "AIS_FLAGS_VALID_COG",
	AIS_FLAGS_VALID_VELOCITY:            "AIS_FLAGS_VALID_VELOCITY",
	AIS_FLAGS_HIGH_VELOCITY:             "AIS_FLAGS_HIGH_VELOCITY",
	AIS_FLAGS_VALID_TURN_RATE:           "AIS_FLAGS_VALID_TURN_RATE",
	AIS_FLAGS_TURN_RATE_SIGN_ONLY:       "AIS_FLAGS_TURN_RATE_SIGN_ONLY",
	AIS_FLAGS_VALID_DIMENSIONS:          "AIS_FLAGS_VALID_DIMENSIONS",
	AIS_FLAGS_LARGE_BOW_DIMENSION:       "AIS_FLAGS_LARGE_BOW_DIMENSION",
	AIS_FLAGS_LARGE_STERN_DIMENSION:     "AIS_FLAGS_LARGE_STERN_DIMENSION",
	AIS_FLAGS_LARGE_PORT_DIMENSION:      "AIS_FLAGS_LARGE_PORT_DIMENSION",
	AIS_FLAGS_LARGE_STARBOARD_DIMENSION: "AIS_FLAGS_LARGE_STARBOARD_DIMENSION",
	AIS_FLAGS_VALID_CALLSIGN:            "AIS_FLAGS_VALID_CALLSIGN",
	AIS_FLAGS_VALID_NAME:                "AIS_FLAGS_VALID_NAME",
}

var values_AIS_FLAGS = map[string]AIS_FLAGS{
	"AIS_FLAGS_POSITION_ACCURACY":         AIS_FLAGS_POSITION_ACCURACY,
	"AIS_FLAGS_VALID_COG":                 AIS_FLAGS_VALID_COG,
	"AIS_FLAGS_VALID_VELOCITY":            AIS_FLAGS_VALID_VELOCITY,
	"AIS_FLAGS_HIGH_VELOCITY":             AIS_FLAGS_HIGH_VELOCITY,
	"AIS_FLAGS_VALID_TURN_RATE":           AIS_FLAGS_VALID_TURN_RATE,
	"AIS_FLAGS_TURN_RATE_SIGN_ONLY":       AIS_FLAGS_TURN_RATE_SIGN_ONLY,
	"AIS_FLAGS_VALID_DIMENSIONS":          AIS_FLAGS_VALID_DIMENSIONS,
	"AIS_FLAGS_LARGE_BOW_DIMENSION":       AIS_FLAGS_LARGE_BOW_DIMENSION,
	"AIS_FLAGS_LARGE_STERN_DIMENSION":     AIS_FLAGS_LARGE_STERN_DIMENSION,
	"AIS_FLAGS_LARGE_PORT_DIMENSION":      AIS_FLAGS_LARGE_PORT_DIMENSION,
	"AIS_FLAGS_LARGE_STARBOARD_DIMENSION": AIS_FLAGS_LARGE_STARBOARD_DIMENSION,
	"AIS_FLAGS_VALID_CALLSIGN":            AIS_FLAGS_VALID_CALLSIGN,
	"AIS_FLAGS_VALID_NAME":                AIS_FLAGS_VALID_NAME,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e AIS_FLAGS) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for i := 0; i < 13; i++ {
		mask := AIS_FLAGS(1 << i)
		if e&mask == mask {
			names = append(names, labels_AIS_FLAGS[mask])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *AIS_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask AIS_FLAGS
	for _, label := range labels {
		if value, ok := values_AIS_FLAGS[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= AIS_FLAGS(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e AIS_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
