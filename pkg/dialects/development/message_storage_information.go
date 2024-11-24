//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Information about a storage medium. This message is sent in response to a request with MAV_CMD_REQUEST_MESSAGE and whenever the status of the storage changes (STORAGE_STATUS). Use MAV_CMD_REQUEST_MESSAGE.param2 to indicate the index/id of requested storage: 0 for all, 1 for first, 2 for second, etc.
type MessageStorageInformation = common.MessageStorageInformation
