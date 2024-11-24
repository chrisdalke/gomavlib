//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

type MAV_ODID_ID_TYPE = common.MAV_ODID_ID_TYPE

const (
	// No type defined.
	MAV_ODID_ID_TYPE_NONE MAV_ODID_ID_TYPE = common.MAV_ODID_ID_TYPE_NONE
	// Manufacturer Serial Number (ANSI/CTA-2063 format).
	MAV_ODID_ID_TYPE_SERIAL_NUMBER MAV_ODID_ID_TYPE = common.MAV_ODID_ID_TYPE_SERIAL_NUMBER
	// CAA (Civil Aviation Authority) registered ID. Format: [ICAO Country Code].[CAA Assigned ID].
	MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID MAV_ODID_ID_TYPE = common.MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID
	// UTM (Unmanned Traffic Management) assigned UUID (RFC4122).
	MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID MAV_ODID_ID_TYPE = common.MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID
	// A 20 byte ID for a specific flight/session. The exact ID type is indicated by the first byte of uas_id and these type values are managed by ICAO.
	MAV_ODID_ID_TYPE_SPECIFIC_SESSION_ID MAV_ODID_ID_TYPE = common.MAV_ODID_ID_TYPE_SPECIFIC_SESSION_ID
)
