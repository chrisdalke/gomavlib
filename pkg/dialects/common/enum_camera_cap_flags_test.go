//autogenerated:yes
//nolint:revive,govet,errcheck
package common

import (
	"testing"
)

func TestEnum_CAMERA_CAP_FLAGS(t *testing.T) {
	var e CAMERA_CAP_FLAGS
	e.UnmarshalText([]byte{})
	e.MarshalText()
	e.String()
}