//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Stream status flags (Bitmap)
type VIDEO_STREAM_STATUS_FLAGS = common.VIDEO_STREAM_STATUS_FLAGS

const (
	// Stream is active (running)
	VIDEO_STREAM_STATUS_FLAGS_RUNNING VIDEO_STREAM_STATUS_FLAGS = common.VIDEO_STREAM_STATUS_FLAGS_RUNNING
	// Stream is thermal imaging
	VIDEO_STREAM_STATUS_FLAGS_THERMAL VIDEO_STREAM_STATUS_FLAGS = common.VIDEO_STREAM_STATUS_FLAGS_THERMAL
	// Stream can report absolute thermal range (see CAMERA_THERMAL_RANGE). (WIP).
	VIDEO_STREAM_STATUS_FLAGS_THERMAL_RANGE_ENABLED VIDEO_STREAM_STATUS_FLAGS = common.VIDEO_STREAM_STATUS_FLAGS_THERMAL_RANGE_ENABLED
)
