//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Terrain data sent from GCS. The lat/lon and grid_spacing must be the same as a lat/lon from a TERRAIN_REQUEST. See terrain protocol docs: https://mavlink.io/en/services/terrain.html
type MessageTerrainData = common.MessageTerrainData
