//autogenerated:yes
//nolint:revive,govet,errcheck
package common

import (
	"testing"
)

func TestEnum_MAV_MISSION_TYPE(t *testing.T) {
	var e MAV_MISSION_TYPE
	e.UnmarshalText([]byte{})
	e.MarshalText()
	e.String()
}