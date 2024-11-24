//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Message reporting the status of a gimbal device.
// This message should be broadcast by a gimbal device component at a low regular rate (e.g. 5 Hz).
// For the angles encoded in the quaternion and the angular velocities holds:
// If the flag GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME is set, then they are relative to the vehicle heading (vehicle frame).
// If the flag GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME is set, then they are relative to absolute North (earth frame).
// If neither of these flags are set, then (for backwards compatibility) it holds:
// If the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set, then they are relative to absolute North (earth frame),
// else they are relative to the vehicle heading (vehicle frame).
// Other conditions of the flags are not allowed.
// The quaternion and angular velocities in the other frame can be calculated from delta_yaw and delta_yaw_velocity as
// q_earth = q_delta_yaw * q_vehicle and w_earth = w_delta_yaw_velocity + w_vehicle (if not NaN).
// If neither the GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME nor the GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME flag is set,
// then (for backwards compatibility) the data in the delta_yaw and delta_yaw_velocity fields are to be ignored.
// New implementations should always set either GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME or GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME,
// and always should set delta_yaw and delta_yaw_velocity either to the proper value or NaN.
type MessageGimbalDeviceAttitudeStatus = common.MessageGimbalDeviceAttitudeStatus
