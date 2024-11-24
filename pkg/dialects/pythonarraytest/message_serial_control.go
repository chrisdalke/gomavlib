//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Control a serial port. This can be used for raw access to an onboard serial peripheral such as a GPS or telemetry radio. It is designed to make it possible to update the devices firmware via MAVLink messages or change the devices settings. A message with zero bytes can be used to change just the baudrate.
type MessageSerialControl = common.MessageSerialControl
