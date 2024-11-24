//autogenerated:yes
//nolint:revive
package dialects

import (
	"encoding"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

var casesEnum = []struct {
	name string
	dec  encoding.TextMarshaler
	enc  string
}{
	{
		"bitmask",
		common.POSITION_TARGET_TYPEMASK_VX_IGNORE | common.POSITION_TARGET_TYPEMASK_VY_IGNORE,
		"POSITION_TARGET_TYPEMASK_VX_IGNORE | POSITION_TARGET_TYPEMASK_VY_IGNORE",
	},
	{
		"value",
		common.GPS_FIX_TYPE_NO_FIX,
		"GPS_FIX_TYPE_NO_FIX",
	},
}

func TestEnumUnmarshalText(t *testing.T) {
	for _, ca := range casesEnum {
		t.Run(ca.name, func(t *testing.T) {
			dec := reflect.New(reflect.TypeOf(ca.dec)).Interface().(encoding.TextUnmarshaler)
			err := dec.UnmarshalText([]byte(ca.enc))
			require.NoError(t, err)
			require.Equal(t, ca.dec, reflect.ValueOf(dec).Elem().Interface())
		})
	}
}

func TestEnumMarshalText(t *testing.T) {
	for _, ca := range casesEnum {
		t.Run(ca.name, func(t *testing.T) {
			byts, err := ca.dec.MarshalText()
			require.NoError(t, err)
			require.Equal(t, ca.enc, string(byts))
		})
	}
}
