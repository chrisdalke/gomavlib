//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/storm32"
)

// Message reporting the current status of a gimbal manager. This message should be broadcast at a low regular rate (e.g. 1 Hz, may be increase momentarily to e.g. 5 Hz for a period of 1 sec after a change).
type MessageStorm32GimbalManagerStatus = storm32.MessageStorm32GimbalManagerStatus
