//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package pythonarraytest

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

type MAV_ODID_AUTH_TYPE = common.MAV_ODID_AUTH_TYPE

const (
	// No authentication type is specified.
	MAV_ODID_AUTH_TYPE_NONE MAV_ODID_AUTH_TYPE = common.MAV_ODID_AUTH_TYPE_NONE
	// Signature for the UAS (Unmanned Aircraft System) ID.
	MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE MAV_ODID_AUTH_TYPE = common.MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE
	// Signature for the Operator ID.
	MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE MAV_ODID_AUTH_TYPE = common.MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE
	// Signature for the entire message set.
	MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE MAV_ODID_AUTH_TYPE = common.MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE
	// Authentication is provided by Network Remote ID.
	MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID MAV_ODID_AUTH_TYPE = common.MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID
	// The exact authentication type is indicated by the first byte of authentication_data and these type values are managed by ICAO.
	MAV_ODID_AUTH_TYPE_SPECIFIC_AUTHENTICATION MAV_ODID_AUTH_TYPE = common.MAV_ODID_AUTH_TYPE_SPECIFIC_AUTHENTICATION
)
