//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/common"
)

// Request a list of available logs. On some systems calling this may stop on-board logging until LOG_REQUEST_END is called. If there are no log files available this request shall be answered with one LOG_ENTRY message with id = 0 and num_logs = 0.
type MessageLogRequestList = common.MessageLogRequestList
