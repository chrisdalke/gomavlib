//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Camera absolute thermal range. This can be streamed when the associated `VIDEO_STREAM_STATUS.flag` bit `VIDEO_STREAM_STATUS_FLAGS_THERMAL_RANGE_ENABLED` is set, but a GCS may choose to only request it for the current active stream. Use MAV_CMD_SET_MESSAGE_INTERVAL to define message interval (param3 indicates the stream id of the current camera, or 0 for all streams, param4 indicates the target camera_device_id for autopilot-attached cameras or 0 for MAVLink cameras).
type MessageCameraThermalRange = common.MessageCameraThermalRange
