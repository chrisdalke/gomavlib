//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It
// is designed as scaled integer message since the resolution of float is not sufficient.
type MessageGlobalPositionInt = common.MessageGlobalPositionInt
