//autogenerated:yes
//nolint:revive
package asluav

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chrisdalke/gomavlib/v3/pkg/dialect"
)

func TestDialect(t *testing.T) {
	_, err := dialect.NewReadWriter(Dialect)
	require.NoError(t, err)
}
