//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Request that the vehicle report terrain height at the given location (expected response is a TERRAIN_REPORT). Used by GCS to check if vehicle has all terrain data needed for a mission.
type MessageTerrainCheck = common.MessageTerrainCheck
