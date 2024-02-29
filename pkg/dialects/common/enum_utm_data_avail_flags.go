//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
	"strings"
)

// Flags for the global position report.
type UTM_DATA_AVAIL_FLAGS uint64

const (
	// The field time contains valid data.
	UTM_DATA_AVAIL_FLAGS_TIME_VALID UTM_DATA_AVAIL_FLAGS = 1
	// The field uas_id contains valid data.
	UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE UTM_DATA_AVAIL_FLAGS = 2
	// The fields lat, lon and h_acc contain valid data.
	UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE UTM_DATA_AVAIL_FLAGS = 4
	// The fields alt and v_acc contain valid data.
	UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 8
	// The field relative_alt contains valid data.
	UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 16
	// The fields vx and vy contain valid data.
	UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 32
	// The field vz contains valid data.
	UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 64
	// The fields next_lat, next_lon and next_alt contain valid data.
	UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE UTM_DATA_AVAIL_FLAGS = 128
)

var labels_UTM_DATA_AVAIL_FLAGS = map[UTM_DATA_AVAIL_FLAGS]string{
	UTM_DATA_AVAIL_FLAGS_TIME_VALID:                  "UTM_DATA_AVAIL_FLAGS_TIME_VALID",
	UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE:            "UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE",
	UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE:          "UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE",
	UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE:          "UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE",
	UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE: "UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE",
	UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE:   "UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE",
	UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE:     "UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE",
	UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE:     "UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE",
}

var values_UTM_DATA_AVAIL_FLAGS = map[string]UTM_DATA_AVAIL_FLAGS{
	"UTM_DATA_AVAIL_FLAGS_TIME_VALID":                  UTM_DATA_AVAIL_FLAGS_TIME_VALID,
	"UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE":            UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE,
	"UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE":          UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE,
	"UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE":          UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE,
	"UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE": UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE,
	"UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE":   UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE,
	"UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE":     UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE,
	"UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE":     UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UTM_DATA_AVAIL_FLAGS) MarshalText() ([]byte, error) {
	if e == 0 {
		return []byte("0"), nil
	}
	var names []string
	for i := 0; i < 8; i++ {
		mask := UTM_DATA_AVAIL_FLAGS(1 << i)
		if e&mask == mask {
			names = append(names, labels_UTM_DATA_AVAIL_FLAGS[mask])
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UTM_DATA_AVAIL_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask UTM_DATA_AVAIL_FLAGS
	for _, label := range labels {
		if value, ok := values_UTM_DATA_AVAIL_FLAGS[label]; ok {
			mask |= value
		} else if value, err := strconv.Atoi(label); err == nil {
			mask |= UTM_DATA_AVAIL_FLAGS(value)
		} else {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e UTM_DATA_AVAIL_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
