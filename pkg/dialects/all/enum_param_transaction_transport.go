//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/development"
)

// Possible transport layers to set and get parameters via mavlink during a parameter transaction.
type PARAM_TRANSACTION_TRANSPORT = development.PARAM_TRANSACTION_TRANSPORT

const (
	// Transaction over param transport.
	PARAM_TRANSACTION_TRANSPORT_PARAM PARAM_TRANSACTION_TRANSPORT = development.PARAM_TRANSACTION_TRANSPORT_PARAM
	// Transaction over param_ext transport.
	PARAM_TRANSACTION_TRANSPORT_PARAM_EXT PARAM_TRANSACTION_TRANSPORT = development.PARAM_TRANSACTION_TRANSPORT_PARAM_EXT
)
