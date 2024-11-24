//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// This message is emitted as response to MISSION_REQUEST_LIST by the MAV and to initiate a write transaction. The GCS can then request the individual mission item based on the knowledge of the total number of waypoints.
type MessageMissionCount = common.MessageMissionCount
