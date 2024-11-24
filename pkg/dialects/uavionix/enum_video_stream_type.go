//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Video stream types
type VIDEO_STREAM_TYPE = common.VIDEO_STREAM_TYPE

const (
	// Stream is RTSP
	VIDEO_STREAM_TYPE_RTSP VIDEO_STREAM_TYPE = common.VIDEO_STREAM_TYPE_RTSP
	// Stream is RTP UDP (URI gives the port number)
	VIDEO_STREAM_TYPE_RTPUDP VIDEO_STREAM_TYPE = common.VIDEO_STREAM_TYPE_RTPUDP
	// Stream is MPEG on TCP
	VIDEO_STREAM_TYPE_TCP_MPEG VIDEO_STREAM_TYPE = common.VIDEO_STREAM_TYPE_TCP_MPEG
	// Stream is MPEG TS (URI gives the port number)
	VIDEO_STREAM_TYPE_MPEG_TS VIDEO_STREAM_TYPE = common.VIDEO_STREAM_TYPE_MPEG_TS
)
