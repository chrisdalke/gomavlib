// Package csairlink contains the csairlink dialect.
//
//autogenerated:yes
package csairlink

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialect"
	"github.com/chrisdalke/gomavlib/v3/pkg/message"
)

// Dialect contains the dialect definition.
var Dialect = dial

// dial is not exposed directly in order not to display it in godoc.
var dial = &dialect.Dialect{
	Version: 3,
	Messages: []message.Message{
		// csairlink
		&MessageAirlinkAuth{},
		&MessageAirlinkAuthResponse{},
		&MessageAirlinkEyeGsHolePushRequest{},
		&MessageAirlinkEyeGsHolePushResponse{},
		&MessageAirlinkEyeHp{},
		&MessageAirlinkEyeTurnInit{},
	},
}
