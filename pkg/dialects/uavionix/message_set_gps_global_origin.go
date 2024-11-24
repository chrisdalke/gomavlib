//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Sets the GPS coordinates of the vehicle local origin (0,0,0) position. Vehicle should emit GPS_GLOBAL_ORIGIN irrespective of whether the origin is changed. This enables transform between the local coordinate frame and the global (GPS) coordinate frame, which may be necessary when (for example) indoor and outdoor settings are connected and the MAV should move from in- to outdoor.
type MessageSetGpsGlobalOrigin = common.MessageSetGpsGlobalOrigin
