//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Message that announces the sequence number of the current target mission item (that the system will fly towards/execute when the mission is running).
// This message should be streamed all the time (nominally at 1Hz).
// This message should be emitted following a call to MAV_CMD_DO_SET_MISSION_CURRENT or SET_MISSION_CURRENT.
type MessageMissionCurrent = common.MessageMissionCurrent
