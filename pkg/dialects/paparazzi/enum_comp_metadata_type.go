//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Supported component metadata types. These are used in the "general" metadata file returned by COMPONENT_METADATA to provide information about supported metadata types. The types are not used directly in MAVLink messages.
type COMP_METADATA_TYPE = common.COMP_METADATA_TYPE

const (
	// General information about the component. General metadata includes information about other metadata types supported by the component. Files of this type must be supported, and must be downloadable from vehicle using a MAVLink FTP URI.
	COMP_METADATA_TYPE_GENERAL COMP_METADATA_TYPE = common.COMP_METADATA_TYPE_GENERAL
	// Parameter meta data.
	COMP_METADATA_TYPE_PARAMETER COMP_METADATA_TYPE = common.COMP_METADATA_TYPE_PARAMETER
	// Meta data that specifies which commands and command parameters the vehicle supports. (WIP)
	COMP_METADATA_TYPE_COMMANDS COMP_METADATA_TYPE = common.COMP_METADATA_TYPE_COMMANDS
	// Meta data that specifies external non-MAVLink peripherals.
	COMP_METADATA_TYPE_PERIPHERALS COMP_METADATA_TYPE = common.COMP_METADATA_TYPE_PERIPHERALS
	// Meta data for the events interface.
	COMP_METADATA_TYPE_EVENTS COMP_METADATA_TYPE = common.COMP_METADATA_TYPE_EVENTS
	// Meta data for actuator configuration (motors, servos and vehicle geometry) and testing.
	COMP_METADATA_TYPE_ACTUATORS COMP_METADATA_TYPE = common.COMP_METADATA_TYPE_ACTUATORS
)
