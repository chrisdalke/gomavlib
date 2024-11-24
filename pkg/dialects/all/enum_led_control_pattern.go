//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/ardupilotmega"
)

type LED_CONTROL_PATTERN = ardupilotmega.LED_CONTROL_PATTERN

const (
	// LED patterns off (return control to regular vehicle control).
	LED_CONTROL_PATTERN_OFF LED_CONTROL_PATTERN = ardupilotmega.LED_CONTROL_PATTERN_OFF
	// LEDs show pattern during firmware update.
	LED_CONTROL_PATTERN_FIRMWAREUPDATE LED_CONTROL_PATTERN = ardupilotmega.LED_CONTROL_PATTERN_FIRMWAREUPDATE
	// Custom Pattern using custom bytes fields.
	LED_CONTROL_PATTERN_CUSTOM LED_CONTROL_PATTERN = ardupilotmega.LED_CONTROL_PATTERN_CUSTOM
)
