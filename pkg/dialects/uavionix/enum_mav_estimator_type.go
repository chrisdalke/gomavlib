//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Enumeration of estimator types
type MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE

const (
	// Unknown type of the estimator.
	MAV_ESTIMATOR_TYPE_UNKNOWN MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_UNKNOWN
	// This is a naive estimator without any real covariance feedback.
	MAV_ESTIMATOR_TYPE_NAIVE MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_NAIVE
	// Computer vision based estimate. Might be up to scale.
	MAV_ESTIMATOR_TYPE_VISION MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_VISION
	// Visual-inertial estimate.
	MAV_ESTIMATOR_TYPE_VIO MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_VIO
	// Plain GPS estimate.
	MAV_ESTIMATOR_TYPE_GPS MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_GPS
	// Estimator integrating GPS and inertial sensing.
	MAV_ESTIMATOR_TYPE_GPS_INS MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_GPS_INS
	// Estimate from external motion capturing system.
	MAV_ESTIMATOR_TYPE_MOCAP MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_MOCAP
	// Estimator based on lidar sensor input.
	MAV_ESTIMATOR_TYPE_LIDAR MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_LIDAR
	// Estimator on autopilot.
	MAV_ESTIMATOR_TYPE_AUTOPILOT MAV_ESTIMATOR_TYPE = common.MAV_ESTIMATOR_TYPE_AUTOPILOT
)
