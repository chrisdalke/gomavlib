//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Battery information that is static, or requires infrequent update.
// This message should requested using MAV_CMD_REQUEST_MESSAGE and/or streamed at very low rate.
// BATTERY_STATUS_V2 is used for higher-rate battery status information.
type MessageBatteryInfo = common.MessageBatteryInfo
