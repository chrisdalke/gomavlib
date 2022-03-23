//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package development

import (
	"errors"
)

// MAV FTP opcodes: https://mavlink.io/en/services/ftp.html
type MAV_FTP_OPCODE uint32

const (
	// None. Ignored, always ACKed
	MAV_FTP_OPCODE_NONE MAV_FTP_OPCODE = 0
	// TerminateSession: Terminates open Read session
	MAV_FTP_OPCODE_TERMINATESESSION MAV_FTP_OPCODE = 1
	// ResetSessions: Terminates all open read sessions
	MAV_FTP_OPCODE_RESETSESSION MAV_FTP_OPCODE = 2
	// ListDirectory. List files and directories in path from offset
	MAV_FTP_OPCODE_LISTDIRECTORY MAV_FTP_OPCODE = 3
	// OpenFileRO: Opens file at path for reading, returns session
	MAV_FTP_OPCODE_OPENFILERO MAV_FTP_OPCODE = 4
	// ReadFile: Reads size bytes from offset in session
	MAV_FTP_OPCODE_READFILE MAV_FTP_OPCODE = 5
	// CreateFile: Creates file at path for writing, returns session
	MAV_FTP_OPCODE_CREATEFILE MAV_FTP_OPCODE = 6
	// WriteFile: Writes size bytes to offset in session
	MAV_FTP_OPCODE_WRITEFILE MAV_FTP_OPCODE = 7
	// RemoveFile: Remove file at path
	MAV_FTP_OPCODE_REMOVEFILE MAV_FTP_OPCODE = 8
	// CreateDirectory: Creates directory at path
	MAV_FTP_OPCODE_CREATEDIRECTORY MAV_FTP_OPCODE = 9
	// RemoveDirectory: Removes directory at path. The directory must be empty.
	MAV_FTP_OPCODE_REMOVEDIRECTORY MAV_FTP_OPCODE = 10
	// OpenFileWO: Opens file at path for writing, returns session
	MAV_FTP_OPCODE_OPENFILEWO MAV_FTP_OPCODE = 11
	// TruncateFile: Truncate file at path to offset length
	MAV_FTP_OPCODE_TRUNCATEFILE MAV_FTP_OPCODE = 12
	// Rename: Rename path1 to path2
	MAV_FTP_OPCODE_RENAME MAV_FTP_OPCODE = 13
	// CalcFileCRC32: Calculate CRC32 for file at path
	MAV_FTP_OPCODE_CALCFILECRC MAV_FTP_OPCODE = 14
	// BurstReadFile: Burst download session file
	MAV_FTP_OPCODE_BURSTREADFILE MAV_FTP_OPCODE = 15
	// ACK: ACK response
	MAV_FTP_OPCODE_ACK MAV_FTP_OPCODE = 128
	// NAK: NAK response
	MAV_FTP_OPCODE_NAK MAV_FTP_OPCODE = 129
)

var labels_MAV_FTP_OPCODE = map[MAV_FTP_OPCODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_FTP_OPCODE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_FTP_OPCODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_FTP_OPCODE = map[string]MAV_FTP_OPCODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_FTP_OPCODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_FTP_OPCODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_FTP_OPCODE) String() string {
	if l, ok := labels_MAV_FTP_OPCODE[e]; ok {
		return l
	}
	return "invalid value"
}
