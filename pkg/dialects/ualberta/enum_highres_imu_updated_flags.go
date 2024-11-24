//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Flags in the HIGHRES_IMU message indicate which fields have updated since the last message
type HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_FLAGS

const (
	// None of the fields in HIGHRES_IMU have been updated
	HIGHRES_IMU_UPDATED_NONE HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_NONE
	// The value in the xacc field has been updated
	HIGHRES_IMU_UPDATED_XACC HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_XACC
	// The value in the yacc field has been updated
	HIGHRES_IMU_UPDATED_YACC HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_YACC
	// The value in the zacc field has been updated since
	HIGHRES_IMU_UPDATED_ZACC HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_ZACC
	// The value in the xgyro field has been updated
	HIGHRES_IMU_UPDATED_XGYRO HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_XGYRO
	// The value in the ygyro field has been updated
	HIGHRES_IMU_UPDATED_YGYRO HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_YGYRO
	// The value in the zgyro field has been updated
	HIGHRES_IMU_UPDATED_ZGYRO HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_ZGYRO
	// The value in the xmag field has been updated
	HIGHRES_IMU_UPDATED_XMAG HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_XMAG
	// The value in the ymag field has been updated
	HIGHRES_IMU_UPDATED_YMAG HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_YMAG
	// The value in the zmag field has been updated
	HIGHRES_IMU_UPDATED_ZMAG HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_ZMAG
	// The value in the abs_pressure field has been updated
	HIGHRES_IMU_UPDATED_ABS_PRESSURE HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_ABS_PRESSURE
	// The value in the diff_pressure field has been updated
	HIGHRES_IMU_UPDATED_DIFF_PRESSURE HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_DIFF_PRESSURE
	// The value in the pressure_alt field has been updated
	HIGHRES_IMU_UPDATED_PRESSURE_ALT HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_PRESSURE_ALT
	// The value in the temperature field has been updated
	HIGHRES_IMU_UPDATED_TEMPERATURE HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_TEMPERATURE
	// All fields in HIGHRES_IMU have been updated.
	HIGHRES_IMU_UPDATED_ALL HIGHRES_IMU_UPDATED_FLAGS = common.HIGHRES_IMU_UPDATED_ALL
)
