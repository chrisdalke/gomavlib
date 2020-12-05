package msg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

type MAV_TYPE int              //nolint:golint
type MAV_AUTOPILOT int         //nolint:golint
type MAV_MODE_FLAG int         //nolint:golint
type MAV_STATE int             //nolint:golint
type MAV_SYS_STATUS_SENSOR int //nolint:golint
type MAV_CMD int               //nolint:golint

type MessageHeartbeat struct {
	Type           MAV_TYPE      `mavenum:"uint8"`
	Autopilot      MAV_AUTOPILOT `mavenum:"uint8"`
	BaseMode       MAV_MODE_FLAG `mavenum:"uint8"`
	CustomMode     uint32
	SystemStatus   MAV_STATE `mavenum:"uint8"`
	MavlinkVersion uint8
}

func (*MessageHeartbeat) GetID() uint32 {
	return 0
}

type MessageSysStatus struct {
	OnboardControlSensorsPresent MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	OnboardControlSensorsEnabled MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	OnboardControlSensorsHealth  MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	Load                         uint16
	VoltageBattery               uint16
	CurrentBattery               int16
	BatteryRemaining             int8
	DropRateComm                 uint16
	ErrorsComm                   uint16
	ErrorsCount1                 uint16
	ErrorsCount2                 uint16
	ErrorsCount3                 uint16
	ErrorsCount4                 uint16
}

func (m *MessageSysStatus) GetID() uint32 {
	return 1
}

type MessageChangeOperatorControl struct {
	TargetSystem   uint8
	ControlRequest uint8
	Version        uint8
	Passkey        string `mavlen:"25"`
}

func (m *MessageChangeOperatorControl) GetID() uint32 {
	return 5
}

type MessageAttitudeQuaternionCov struct {
	TimeUsec   uint64
	Q          [4]float32
	Rollspeed  float32
	Pitchspeed float32
	Yawspeed   float32
	Covariance [9]float32
}

func (m *MessageAttitudeQuaternionCov) GetID() uint32 {
	return 61
}

type MessageOpticalFlow struct {
	TimeUsec       uint64
	SensorId       uint8 //nolint:golint
	FlowX          int16
	FlowY          int16
	FlowCompMX     float32
	FlowCompMY     float32
	Quality        uint8
	GroundDistance float32
	FlowRateX      float32 `mavext:"true"`
	FlowRateY      float32 `mavext:"true"`
}

func (*MessageOpticalFlow) GetID() uint32 {
	return 100
}

type MessagePlayTune struct {
	TargetSystem    uint8
	TargetComponent uint8
	Tune            string `mavlen:"30"`
	Tune2           string `mavext:"true" mavlen:"200"`
}

func (*MessagePlayTune) GetID() uint32 {
	return 258
}

type MessageAhrs struct {
	OmegaIx     float32 `mavname:"omegaIx"`
	OmegaIy     float32 `mavname:"omegaIy"`
	OmegaIz     float32 `mavname:"omegaIz"`
	AccelWeight float32
	RenormVal   float32
	ErrorRp     float32
	ErrorYaw    float32
}

func (*MessageAhrs) GetID() uint32 {
	return 163
}

type MessageTrajectoryRepresentationWaypoints struct {
	TimeUsec    uint64
	ValidPoints uint8
	PosX        [5]float32
	PosY        [5]float32
	PosZ        [5]float32
	VelX        [5]float32
	VelY        [5]float32
	VelZ        [5]float32
	AccX        [5]float32
	AccY        [5]float32
	AccZ        [5]float32
	PosYaw      [5]float32
	VelYaw      [5]float32
	Command     [5]MAV_CMD `mavenum:"uint16"`
}

func (*MessageTrajectoryRepresentationWaypoints) GetID() uint32 {
	return 332
}

var casesCRC = []struct {
	msg Message
	crc byte
}{
	{
		&MessageHeartbeat{},
		50,
	},
	{
		&MessageSysStatus{},
		124,
	},
	{
		&MessageChangeOperatorControl{},
		217,
	},
	{
		&MessageAttitudeQuaternionCov{},
		167,
	},
	{
		&MessageOpticalFlow{},
		175,
	},
	{
		&MessagePlayTune{},
		187,
	},
	{
		&MessageAhrs{},
		127,
	},
}

func TestCRC(t *testing.T) {
	for _, c := range casesCRC {
		mp, err := NewDecEncoder(c.msg)
		require.NoError(t, err)
		require.Equal(t, c.crc, mp.crcExtra)
	}
}

var casesMsgs = []struct {
	name   string
	isV2   bool
	parsed Message
	raw    []byte
}{
	{
		"v1 basic a",
		false,
		&MessageHeartbeat{
			Type:           1,
			Autopilot:      2,
			BaseMode:       3,
			CustomMode:     6,
			SystemStatus:   4,
			MavlinkVersion: 5,
		},
		[]byte("\x06\x00\x00\x00\x01\x02\x03\x04\x05"),
	},
	{
		"v1 basic b",
		false,
		&MessageSysStatus{
			OnboardControlSensorsPresent: 0x01010101,
			OnboardControlSensorsEnabled: 0x01010101,
			OnboardControlSensorsHealth:  0x01010101,
			Load:                         0x0101,
			VoltageBattery:               0x0101,
			CurrentBattery:               0x0101,
			BatteryRemaining:             1,
			DropRateComm:                 0x0101,
			ErrorsComm:                   0x0101,
			ErrorsCount1:                 0x0101,
			ErrorsCount2:                 0x0101,
			ErrorsCount3:                 0x0101,
			ErrorsCount4:                 0x0101,
		},
		bytes.Repeat([]byte("\x01"), 31),
	},
	{
		"v1 basic c",
		false,
		&MessageChangeOperatorControl{
			TargetSystem:   1,
			ControlRequest: 1,
			Version:        1,
			Passkey:        "testing",
		},
		[]byte("\x01\x01\x01\x74\x65\x73\x74\x69\x6e\x67\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
	},
	{
		"v1 with string with max length",
		false,
		&MessageChangeOperatorControl{
			Passkey: "abcdefghijklmnopqrstuvwxy",
		},
		[]byte("\x00\x00\x00\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70\x71\x72\x73\x74\x75\x76\x77\x78\x79"),
	},
	{
		"v1 with array",
		false,
		&MessageAttitudeQuaternionCov{
			TimeUsec:   2,
			Q:          [4]float32{1, 1, 1, 1},
			Rollspeed:  1,
			Pitchspeed: 1,
			Yawspeed:   1,
			Covariance: [9]float32{1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		append([]byte("\x02\x00\x00\x00\x00\x00\x00\x00"), bytes.Repeat([]byte("\x00\x00\x80\x3F"), 16)...),
	},
	{
		"v1 with extensions",
		false,
		&MessageOpticalFlow{
			TimeUsec:       3,
			FlowCompMX:     1,
			FlowCompMY:     1,
			GroundDistance: 1,
			FlowX:          7,
			FlowY:          8,
			SensorId:       9,
			Quality:        0x0A,
		},
		[]byte("\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\x3F\x00\x00\x80\x3F\x00\x00\x80\x3F\x07\x00\x08\x00\x09\x0A"),
	},
	{
		"v1 with array of enums",
		false,
		&MessageTrajectoryRepresentationWaypoints{
			TimeUsec:    1,
			ValidPoints: 2,
			PosX:        [5]float32{1, 2, 3, 4, 5},
			PosY:        [5]float32{1, 2, 3, 4, 5},
			PosZ:        [5]float32{1, 2, 3, 4, 5},
			VelX:        [5]float32{1, 2, 3, 4, 5},
			VelY:        [5]float32{1, 2, 3, 4, 5},
			VelZ:        [5]float32{1, 2, 3, 4, 5},
			AccX:        [5]float32{1, 2, 3, 4, 5},
			AccY:        [5]float32{1, 2, 3, 4, 5},
			AccZ:        [5]float32{1, 2, 3, 4, 5},
			PosYaw:      [5]float32{1, 2, 3, 4, 5},
			VelYaw:      [5]float32{1, 2, 3, 4, 5},
			Command:     [5]MAV_CMD{1, 2, 3, 4, 5},
		},
		[]byte("\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x01\x00\x02\x00\x03\x00\x04\x00\x05\x00\x02"),
	},
	{
		"v2 with empty-byte truncation and empty",
		true,
		&MessageAhrs{},
		[]byte("\x00"),
	},
	{
		"v2 with empty-byte truncation a",
		true,
		&MessageChangeOperatorControl{
			TargetSystem:   0,
			ControlRequest: 1,
			Version:        2,
			Passkey:        "testing",
		},
		[]byte("\x00\x01\x02\x74\x65\x73\x74\x69\x6e\x67"),
	},
	{
		"v2 with empty-byte truncation b",
		true,
		&MessageAhrs{
			OmegaIx:     1,
			OmegaIy:     2,
			OmegaIz:     3,
			AccelWeight: 4,
			RenormVal:   5,
			ErrorRp:     0,
			ErrorYaw:    0,
		},
		[]byte("\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40"),
	},
	{
		"v2 with extensions a",
		true,
		&MessageOpticalFlow{
			TimeUsec:       3,
			FlowCompMX:     1,
			FlowCompMY:     1,
			GroundDistance: 1,
			FlowX:          7,
			FlowY:          8,
			SensorId:       9,
			Quality:        0x0A,
			FlowRateX:      1,
			FlowRateY:      1,
		},
		[]byte("\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\x3F\x00\x00\x80\x3F\x00\x00\x80\x3F\x07\x00\x08\x00\x09\x0A\x00\x00\x80\x3F\x00\x00\x80\x3F"),
	},
	{
		"v2 with extensions b",
		true,
		&MessagePlayTune{
			TargetSystem:    1,
			TargetComponent: 2,
			Tune:            "test1",
			Tune2:           "test2",
		},
		[]byte("\x01\x02\x74\x65\x73\x74\x31\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x74\x65\x73\x74\x32"),
	},
}

func TestDecode(t *testing.T) {
	for _, c := range casesMsgs {
		t.Run(c.name, func(t *testing.T) {
			mp, err := NewDecEncoder(c.parsed)
			require.NoError(t, err)
			msg, err := mp.Decode(c.raw, c.isV2)
			require.NoError(t, err)
			require.Equal(t, c.parsed, msg)
		})
	}
}

func TestEncode(t *testing.T) {
	for _, c := range casesMsgs {
		t.Run(c.name, func(t *testing.T) {
			mp, err := NewDecEncoder(c.parsed)
			require.NoError(t, err)
			byt, err := mp.Encode(c.parsed, c.isV2)
			require.NoError(t, err)
			require.Equal(t, c.raw, byt)
		})
	}
}
