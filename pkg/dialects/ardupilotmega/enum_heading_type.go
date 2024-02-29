//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strconv"
)

type HEADING_TYPE uint64

const (
	HEADING_TYPE_COURSE_OVER_GROUND HEADING_TYPE = 0
	HEADING_TYPE_HEADING            HEADING_TYPE = 1
)

var labels_HEADING_TYPE = map[HEADING_TYPE]string{
	HEADING_TYPE_COURSE_OVER_GROUND: "HEADING_TYPE_COURSE_OVER_GROUND",
	HEADING_TYPE_HEADING:            "HEADING_TYPE_HEADING",
}

var values_HEADING_TYPE = map[string]HEADING_TYPE{
	"HEADING_TYPE_COURSE_OVER_GROUND": HEADING_TYPE_COURSE_OVER_GROUND,
	"HEADING_TYPE_HEADING":            HEADING_TYPE_HEADING,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e HEADING_TYPE) MarshalText() ([]byte, error) {
	if name, ok := labels_HEADING_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *HEADING_TYPE) UnmarshalText(text []byte) error {
	if value, ok := values_HEADING_TYPE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = HEADING_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e HEADING_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
