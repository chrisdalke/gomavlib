//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Specifies the datatype of a MAVLink extended parameter.
type MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE

const (
	// 8-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT8 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_UINT8
	// 8-bit signed integer
	MAV_PARAM_EXT_TYPE_INT8 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_INT8
	// 16-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT16 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_UINT16
	// 16-bit signed integer
	MAV_PARAM_EXT_TYPE_INT16 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_INT16
	// 32-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT32 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_UINT32
	// 32-bit signed integer
	MAV_PARAM_EXT_TYPE_INT32 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_INT32
	// 64-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT64 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_UINT64
	// 64-bit signed integer
	MAV_PARAM_EXT_TYPE_INT64 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_INT64
	// 32-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL32 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_REAL32
	// 64-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL64 MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_REAL64
	// Custom Type
	MAV_PARAM_EXT_TYPE_CUSTOM MAV_PARAM_EXT_TYPE = common.MAV_PARAM_EXT_TYPE_CUSTOM
)
