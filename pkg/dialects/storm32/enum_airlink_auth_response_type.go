//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/csairlink"
)

type AIRLINK_AUTH_RESPONSE_TYPE = csairlink.AIRLINK_AUTH_RESPONSE_TYPE

const (
	// Login or password error
	AIRLINK_ERROR_LOGIN_OR_PASS AIRLINK_AUTH_RESPONSE_TYPE = csairlink.AIRLINK_ERROR_LOGIN_OR_PASS
	// Auth successful
	AIRLINK_AUTH_OK AIRLINK_AUTH_RESPONSE_TYPE = csairlink.AIRLINK_AUTH_OK
)
