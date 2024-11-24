//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Request a partial list of mission items from the system/component. https://mavlink.io/en/services/mission.html. If start and end index are the same, just send one waypoint.
type MessageMissionRequestPartialList = common.MessageMissionRequestPartialList
