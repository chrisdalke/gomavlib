//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/ardupilotmega"
)

type GOPRO_REQUEST_STATUS = ardupilotmega.GOPRO_REQUEST_STATUS

const (
	// The write message with ID indicated succeeded.
	GOPRO_REQUEST_SUCCESS GOPRO_REQUEST_STATUS = ardupilotmega.GOPRO_REQUEST_SUCCESS
	// The write message with ID indicated failed.
	GOPRO_REQUEST_FAILED GOPRO_REQUEST_STATUS = ardupilotmega.GOPRO_REQUEST_FAILED
)
