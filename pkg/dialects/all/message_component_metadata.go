//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Component metadata message, which may be requested using MAV_CMD_REQUEST_MESSAGE.
// This contains the MAVLink FTP URI and CRC for the component's general metadata file.
// The file must be hosted on the component, and may be xz compressed.
// The file CRC can be used for file caching.
// The general metadata file can be read to get the locations of other metadata files (COMP_METADATA_TYPE) and translations, which may be hosted either on the vehicle or the internet.
// For more information see: https://mavlink.io/en/services/component_information.html.
// Note: Camera components should use CAMERA_INFORMATION instead, and autopilots may use both this message and AUTOPILOT_VERSION.
type MessageComponentMetadata = common.MessageComponentMetadata
