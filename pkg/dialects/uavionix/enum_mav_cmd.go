//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/common"
)

// Commands to be executed by the MAV. They can be executed on user request, or as part of a mission script. If the action is used in a mission, the parameter mapping to the waypoint/mission message is as follows: Param 1, Param 2, Param 3, Param 4, X: Param 5, Y:Param 6, Z:Param 7. This command list is similar what ARINC 424 is for commercial aircraft: A data format how to interpret waypoint/mission data. NaN and INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current yaw or latitude rather than a specific value). See https://mavlink.io/en/guide/xml_schema.html#MAV_CMD for information about the structure of the MAV_CMD entries
type MAV_CMD = common.MAV_CMD

const (
	// Navigate to waypoint. This is intended for use in missions (for guided commands outside of missions use MAV_CMD_DO_REPOSITION).
	MAV_CMD_NAV_WAYPOINT MAV_CMD = common.MAV_CMD_NAV_WAYPOINT
	// Loiter around this waypoint an unlimited amount of time
	MAV_CMD_NAV_LOITER_UNLIM MAV_CMD = common.MAV_CMD_NAV_LOITER_UNLIM
	// Loiter around this waypoint for X turns
	MAV_CMD_NAV_LOITER_TURNS MAV_CMD = common.MAV_CMD_NAV_LOITER_TURNS
	// Loiter at the specified latitude, longitude and altitude for a certain amount of time. Multicopter vehicles stop at the point (within a vehicle-specific acceptance radius). Forward-only moving vehicles (e.g. fixed-wing) circle the point with the specified radius/direction. If the Heading Required parameter (2) is non-zero forward moving aircraft will only leave the loiter circle once heading towards the next waypoint.
	MAV_CMD_NAV_LOITER_TIME MAV_CMD = common.MAV_CMD_NAV_LOITER_TIME
	// Return to launch location
	MAV_CMD_NAV_RETURN_TO_LAUNCH MAV_CMD = common.MAV_CMD_NAV_RETURN_TO_LAUNCH
	// Land at location.
	MAV_CMD_NAV_LAND MAV_CMD = common.MAV_CMD_NAV_LAND
	// Takeoff from ground / hand. Vehicles that support multiple takeoff modes (e.g. VTOL quadplane) should take off using the currently configured mode.
	MAV_CMD_NAV_TAKEOFF MAV_CMD = common.MAV_CMD_NAV_TAKEOFF
	// Land at local position (local frame only)
	MAV_CMD_NAV_LAND_LOCAL MAV_CMD = common.MAV_CMD_NAV_LAND_LOCAL
	// Takeoff from local position (local frame only)
	MAV_CMD_NAV_TAKEOFF_LOCAL MAV_CMD = common.MAV_CMD_NAV_TAKEOFF_LOCAL
	// Vehicle following, i.e. this waypoint represents the position of a moving vehicle
	MAV_CMD_NAV_FOLLOW MAV_CMD = common.MAV_CMD_NAV_FOLLOW
	// Continue on the current course and climb/descend to specified altitude.  When the altitude is reached continue to the next command (i.e., don't proceed to the next command until the desired altitude is reached.
	MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT MAV_CMD = common.MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT
	// Begin loiter at the specified Latitude and Longitude.  If Lat=Lon=0, then loiter at the current position.  Don't consider the navigation command complete (don't leave loiter) until the altitude has been reached. Additionally, if the Heading Required parameter is non-zero the aircraft will not leave the loiter until heading toward the next waypoint.
	MAV_CMD_NAV_LOITER_TO_ALT MAV_CMD = common.MAV_CMD_NAV_LOITER_TO_ALT
	// Begin following a target
	MAV_CMD_DO_FOLLOW MAV_CMD = common.MAV_CMD_DO_FOLLOW
	// Reposition the MAV after a follow target command has been sent
	MAV_CMD_DO_FOLLOW_REPOSITION MAV_CMD = common.MAV_CMD_DO_FOLLOW_REPOSITION
	// Start orbiting on the circumference of a circle defined by the parameters. Setting values to NaN/INT32_MAX (as appropriate) results in using defaults.
	MAV_CMD_DO_ORBIT MAV_CMD = common.MAV_CMD_DO_ORBIT
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	MAV_CMD_NAV_ROI MAV_CMD = common.MAV_CMD_NAV_ROI
	// Control autonomous path planning on the MAV.
	MAV_CMD_NAV_PATHPLANNING MAV_CMD = common.MAV_CMD_NAV_PATHPLANNING
	// Navigate to waypoint using a spline path.
	MAV_CMD_NAV_SPLINE_WAYPOINT MAV_CMD = common.MAV_CMD_NAV_SPLINE_WAYPOINT
	// Takeoff from ground using VTOL mode, and transition to forward flight with specified heading. The command should be ignored by vehicles that dont support both VTOL and fixed-wing flight (multicopters, boats,etc.).
	MAV_CMD_NAV_VTOL_TAKEOFF MAV_CMD = common.MAV_CMD_NAV_VTOL_TAKEOFF
	// Land using VTOL mode
	MAV_CMD_NAV_VTOL_LAND MAV_CMD = common.MAV_CMD_NAV_VTOL_LAND
	// hand control over to an external controller
	MAV_CMD_NAV_GUIDED_ENABLE MAV_CMD = common.MAV_CMD_NAV_GUIDED_ENABLE
	// Delay the next navigation command a number of seconds or until a specified time
	MAV_CMD_NAV_DELAY MAV_CMD = common.MAV_CMD_NAV_DELAY
	// Descend and place payload. Vehicle moves to specified location, descends until it detects a hanging payload has reached the ground, and then releases the payload. If ground is not detected before the reaching the maximum descent value (param1), the command will complete without releasing the payload.
	MAV_CMD_NAV_PAYLOAD_PLACE MAV_CMD = common.MAV_CMD_NAV_PAYLOAD_PLACE
	// NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration
	MAV_CMD_NAV_LAST MAV_CMD = common.MAV_CMD_NAV_LAST
	// Delay mission state machine.
	MAV_CMD_CONDITION_DELAY MAV_CMD = common.MAV_CMD_CONDITION_DELAY
	// Ascend/descend to target altitude at specified rate. Delay mission state machine until desired altitude reached.
	MAV_CMD_CONDITION_CHANGE_ALT MAV_CMD = common.MAV_CMD_CONDITION_CHANGE_ALT
	// Delay mission state machine until within desired distance of next NAV point.
	MAV_CMD_CONDITION_DISTANCE MAV_CMD = common.MAV_CMD_CONDITION_DISTANCE
	// Reach a certain target angle.
	MAV_CMD_CONDITION_YAW MAV_CMD = common.MAV_CMD_CONDITION_YAW
	// NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration
	MAV_CMD_CONDITION_LAST MAV_CMD = common.MAV_CMD_CONDITION_LAST
	// Set system mode.
	MAV_CMD_DO_SET_MODE MAV_CMD = common.MAV_CMD_DO_SET_MODE
	// Jump to the desired command in the mission list.  Repeat this action only the specified number of times
	MAV_CMD_DO_JUMP MAV_CMD = common.MAV_CMD_DO_JUMP
	// Change speed and/or throttle set points. The value persists until it is overridden or there is a mode change.
	MAV_CMD_DO_CHANGE_SPEED MAV_CMD = common.MAV_CMD_DO_CHANGE_SPEED
	// Sets the home position to either to the current position or a specified position.
	// The home position is the default position that the system will return to and land on.
	// The position is set automatically by the system during the takeoff (and may also be set using this command).
	// Note: the current home position may be emitted in a HOME_POSITION message on request (using MAV_CMD_REQUEST_MESSAGE with param1=242).
	MAV_CMD_DO_SET_HOME MAV_CMD = common.MAV_CMD_DO_SET_HOME
	// Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter.
	MAV_CMD_DO_SET_PARAMETER MAV_CMD = common.MAV_CMD_DO_SET_PARAMETER
	// Set a relay to a condition.
	MAV_CMD_DO_SET_RELAY MAV_CMD = common.MAV_CMD_DO_SET_RELAY
	// Cycle a relay on and off for a desired number of cycles with a desired period.
	MAV_CMD_DO_REPEAT_RELAY MAV_CMD = common.MAV_CMD_DO_REPEAT_RELAY
	// Set a servo to a desired PWM value.
	MAV_CMD_DO_SET_SERVO MAV_CMD = common.MAV_CMD_DO_SET_SERVO
	// Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period.
	MAV_CMD_DO_REPEAT_SERVO MAV_CMD = common.MAV_CMD_DO_REPEAT_SERVO
	// Terminate flight immediately.
	// Flight termination immediately and irreversably terminates the current flight, returning the vehicle to ground.
	// The vehicle will ignore RC or other input until it has been power-cycled.
	// Termination may trigger safety measures, including: disabling motors and deployment of parachute on multicopters, and setting flight surfaces to initiate a landing pattern on fixed-wing).
	// On multicopters without a parachute it may trigger a crash landing.
	// Support for this command can be tested using the protocol bit: MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION.
	// Support for this command can also be tested by sending the command with param1=0 (< 0.5); the ACK should be either MAV_RESULT_FAILED or MAV_RESULT_UNSUPPORTED.
	MAV_CMD_DO_FLIGHTTERMINATION MAV_CMD = common.MAV_CMD_DO_FLIGHTTERMINATION
	// Change altitude set point.
	MAV_CMD_DO_CHANGE_ALTITUDE MAV_CMD = common.MAV_CMD_DO_CHANGE_ALTITUDE
	// Sets actuators (e.g. servos) to a desired value. The actuator numbers are mapped to specific outputs (e.g. on any MAIN or AUX PWM or UAVCAN) using a flight-stack specific mechanism (i.e. a parameter).
	MAV_CMD_DO_SET_ACTUATOR MAV_CMD = common.MAV_CMD_DO_SET_ACTUATOR
	// Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts.
	// It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used.
	// The Latitude/Longitude/Altitude is optional, and may be set to 0 if not needed. If specified then it will be used to help find the closest landing sequence.
	MAV_CMD_DO_LAND_START MAV_CMD = common.MAV_CMD_DO_LAND_START
	// Mission command to perform a landing from a rally point.
	MAV_CMD_DO_RALLY_LAND MAV_CMD = common.MAV_CMD_DO_RALLY_LAND
	// Mission command to safely abort an autonomous landing.
	MAV_CMD_DO_GO_AROUND MAV_CMD = common.MAV_CMD_DO_GO_AROUND
	// Reposition the vehicle to a specific WGS84 global position. This command is intended for guided commands (for missions use MAV_CMD_NAV_WAYPOINT instead).
	MAV_CMD_DO_REPOSITION MAV_CMD = common.MAV_CMD_DO_REPOSITION
	// If in a GPS controlled position mode, hold the current position or continue.
	MAV_CMD_DO_PAUSE_CONTINUE MAV_CMD = common.MAV_CMD_DO_PAUSE_CONTINUE
	// Set moving direction to forward or reverse.
	MAV_CMD_DO_SET_REVERSE MAV_CMD = common.MAV_CMD_DO_SET_REVERSE
	// Sets the region of interest (ROI) to a location. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal is not to react to this message.
	MAV_CMD_DO_SET_ROI_LOCATION MAV_CMD = common.MAV_CMD_DO_SET_ROI_LOCATION
	// Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message.
	MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET MAV_CMD = common.MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET
	// Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message. After this command the gimbal manager should go back to manual input if available, and otherwise assume a neutral position.
	MAV_CMD_DO_SET_ROI_NONE MAV_CMD = common.MAV_CMD_DO_SET_ROI_NONE
	// Mount tracks system with specified system ID. Determination of target vehicle position may be done with GLOBAL_POSITION_INT or any other means. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message.
	MAV_CMD_DO_SET_ROI_SYSID MAV_CMD = common.MAV_CMD_DO_SET_ROI_SYSID
	// Control onboard camera system.
	MAV_CMD_DO_CONTROL_VIDEO MAV_CMD = common.MAV_CMD_DO_CONTROL_VIDEO
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	MAV_CMD_DO_SET_ROI MAV_CMD = common.MAV_CMD_DO_SET_ROI
	// Configure digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	MAV_CMD_DO_DIGICAM_CONFIGURE MAV_CMD = common.MAV_CMD_DO_DIGICAM_CONFIGURE
	// Control digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	MAV_CMD_DO_DIGICAM_CONTROL MAV_CMD = common.MAV_CMD_DO_DIGICAM_CONTROL
	// Mission command to configure a camera or antenna mount
	MAV_CMD_DO_MOUNT_CONFIGURE MAV_CMD = common.MAV_CMD_DO_MOUNT_CONFIGURE
	// Mission command to control a camera or antenna mount
	MAV_CMD_DO_MOUNT_CONTROL MAV_CMD = common.MAV_CMD_DO_MOUNT_CONTROL
	// Mission command to set camera trigger distance for this flight. The camera is triggered each time this distance is exceeded. This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_DO_SET_CAM_TRIGG_DIST MAV_CMD = common.MAV_CMD_DO_SET_CAM_TRIGG_DIST
	// Mission command to enable the geofence
	MAV_CMD_DO_FENCE_ENABLE MAV_CMD = common.MAV_CMD_DO_FENCE_ENABLE
	// Mission item/command to release a parachute or enable/disable auto release.
	MAV_CMD_DO_PARACHUTE MAV_CMD = common.MAV_CMD_DO_PARACHUTE
	// Command to perform motor test.
	MAV_CMD_DO_MOTOR_TEST MAV_CMD = common.MAV_CMD_DO_MOTOR_TEST
	// Change to/from inverted flight.
	MAV_CMD_DO_INVERTED_FLIGHT MAV_CMD = common.MAV_CMD_DO_INVERTED_FLIGHT
	// Mission command to operate a gripper.
	MAV_CMD_DO_GRIPPER MAV_CMD = common.MAV_CMD_DO_GRIPPER
	// Enable/disable autotune.
	MAV_CMD_DO_AUTOTUNE_ENABLE MAV_CMD = common.MAV_CMD_DO_AUTOTUNE_ENABLE
	// Sets a desired vehicle turn angle and speed change.
	MAV_CMD_NAV_SET_YAW_SPEED MAV_CMD = common.MAV_CMD_NAV_SET_YAW_SPEED
	// Mission command to set camera trigger interval for this flight. If triggering is enabled, the camera is triggered each time this interval expires. This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL MAV_CMD = common.MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL
	// Mission command to control a camera or antenna mount, using a quaternion as reference.
	MAV_CMD_DO_MOUNT_CONTROL_QUAT MAV_CMD = common.MAV_CMD_DO_MOUNT_CONTROL_QUAT
	// set id of master controller
	MAV_CMD_DO_GUIDED_MASTER MAV_CMD = common.MAV_CMD_DO_GUIDED_MASTER
	// Set limits for external control
	MAV_CMD_DO_GUIDED_LIMITS MAV_CMD = common.MAV_CMD_DO_GUIDED_LIMITS
	// Control vehicle engine. This is interpreted by the vehicles engine controller to change the target engine state. It is intended for vehicles with internal combustion engines
	MAV_CMD_DO_ENGINE_CONTROL MAV_CMD = common.MAV_CMD_DO_ENGINE_CONTROL
	// Set the mission item with sequence number seq as the current item and emit MISSION_CURRENT (whether or not the mission number changed).
	// If a mission is currently being executed, the system will continue to this new mission item on the shortest path, skipping any intermediate mission items.
	// Note that mission jump repeat counters are not reset unless param2 is set (see MAV_CMD_DO_JUMP param2).
	// This command may trigger a mission state-machine change on some systems: for example from MISSION_STATE_NOT_STARTED or MISSION_STATE_PAUSED to MISSION_STATE_ACTIVE.
	// If the system is in mission mode, on those systems this command might therefore start, restart or resume the mission.
	// If the system is not in mission mode this command must not trigger a switch to mission mode.
	// The mission may be "reset" using param2.
	// Resetting sets jump counters to initial values (to reset counters without changing the current mission item set the param1 to `-1`).
	// Resetting also explicitly changes a mission state of MISSION_STATE_COMPLETE to MISSION_STATE_PAUSED or MISSION_STATE_ACTIVE, potentially allowing it to resume when it is (next) in a mission mode.
	// The command will ACK with MAV_RESULT_FAILED if the sequence number is out of range (including if there is no mission item).
	MAV_CMD_DO_SET_MISSION_CURRENT MAV_CMD = common.MAV_CMD_DO_SET_MISSION_CURRENT
	// NOP - This command is only used to mark the upper limit of the DO commands in the enumeration
	MAV_CMD_DO_LAST MAV_CMD = common.MAV_CMD_DO_LAST
	// Trigger calibration. This command will be only accepted if in pre-flight mode. Except for Temperature Calibration, only one sensor should be set in a single message and all others should be zero.
	MAV_CMD_PREFLIGHT_CALIBRATION MAV_CMD = common.MAV_CMD_PREFLIGHT_CALIBRATION
	// Set sensor offsets. This command will be only accepted if in pre-flight mode.
	MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS MAV_CMD = common.MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS
	// Trigger UAVCAN configuration (actuator ID assignment and direction mapping). Note that this maps to the legacy UAVCAN v0 function UAVCAN_ENUMERATE, which is intended to be executed just once during initial vehicle configuration (it is not a normal pre-flight command and has been poorly named).
	MAV_CMD_PREFLIGHT_UAVCAN MAV_CMD = common.MAV_CMD_PREFLIGHT_UAVCAN
	// Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode.
	MAV_CMD_PREFLIGHT_STORAGE MAV_CMD = common.MAV_CMD_PREFLIGHT_STORAGE
	// Request the reboot or shutdown of system components.
	MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN MAV_CMD = common.MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN
	// Override current mission with command to pause mission, pause mission and move to position, continue/resume mission. When param 1 indicates that the mission is paused (MAV_GOTO_DO_HOLD), param 2 defines whether it holds in place or moves to another position.
	MAV_CMD_OVERRIDE_GOTO MAV_CMD = common.MAV_CMD_OVERRIDE_GOTO
	// Mission command to set a Camera Auto Mount Pivoting Oblique Survey (Replaces CAM_TRIGG_DIST for this purpose). The camera is triggered each time this distance is exceeded, then the mount moves to the next position. Params 4~6 set-up the angle limits and number of positions for oblique survey, where mount-enabled vehicles automatically roll the camera between shots to emulate an oblique camera setup (providing an increased HFOV). This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_OBLIQUE_SURVEY MAV_CMD = common.MAV_CMD_OBLIQUE_SURVEY
	// start running a mission
	MAV_CMD_MISSION_START MAV_CMD = common.MAV_CMD_MISSION_START
	// Actuator testing command. This is similar to MAV_CMD_DO_MOTOR_TEST but operates on the level of output functions, i.e. it is possible to test Motor1 independent from which output it is configured on. Autopilots typically refuse this command while armed.
	MAV_CMD_ACTUATOR_TEST MAV_CMD = common.MAV_CMD_ACTUATOR_TEST
	// Actuator configuration command.
	MAV_CMD_CONFIGURE_ACTUATOR MAV_CMD = common.MAV_CMD_CONFIGURE_ACTUATOR
	// Arms / Disarms a component
	MAV_CMD_COMPONENT_ARM_DISARM MAV_CMD = common.MAV_CMD_COMPONENT_ARM_DISARM
	// Instructs a target system to run pre-arm checks.
	// This allows preflight checks to be run on demand, which may be useful on systems that normally run them at low rate, or which do not trigger checks when the armable state might have changed.
	// This command should return MAV_RESULT_ACCEPTED if it will run the checks.
	// The results of the checks are usually then reported in SYS_STATUS messages (this is system-specific).
	// The command should return MAV_RESULT_TEMPORARILY_REJECTED if the system is already armed.
	MAV_CMD_RUN_PREARM_CHECKS MAV_CMD = common.MAV_CMD_RUN_PREARM_CHECKS
	// Turns illuminators ON/OFF. An illuminator is a light source that is used for lighting up dark areas external to the sytstem: e.g. a torch or searchlight (as opposed to a light source for illuminating the system itself, e.g. an indicator light).
	MAV_CMD_ILLUMINATOR_ON_OFF MAV_CMD = common.MAV_CMD_ILLUMINATOR_ON_OFF
	// Request the home position from the vehicle.
	// The vehicle will ACK the command and then emit the HOME_POSITION message.
	MAV_CMD_GET_HOME_POSITION MAV_CMD = common.MAV_CMD_GET_HOME_POSITION
	// Inject artificial failure for testing purposes. Note that autopilots should implement an additional protection before accepting this command such as a specific param setting.
	MAV_CMD_INJECT_FAILURE MAV_CMD = common.MAV_CMD_INJECT_FAILURE
	// Starts receiver pairing.
	MAV_CMD_START_RX_PAIR MAV_CMD = common.MAV_CMD_START_RX_PAIR
	// Request the interval between messages for a particular MAVLink message ID.
	// The receiver should ACK the command and then emit its response in a MESSAGE_INTERVAL message.
	MAV_CMD_GET_MESSAGE_INTERVAL MAV_CMD = common.MAV_CMD_GET_MESSAGE_INTERVAL
	// Set the interval between messages for a particular MAVLink message ID. This interface replaces REQUEST_DATA_STREAM.
	MAV_CMD_SET_MESSAGE_INTERVAL MAV_CMD = common.MAV_CMD_SET_MESSAGE_INTERVAL
	// Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of MAV_CMD_SET_MESSAGE_INTERVAL).
	MAV_CMD_REQUEST_MESSAGE MAV_CMD = common.MAV_CMD_REQUEST_MESSAGE
	// Request MAVLink protocol version compatibility. All receivers should ACK the command and then emit their capabilities in an PROTOCOL_VERSION message
	MAV_CMD_REQUEST_PROTOCOL_VERSION MAV_CMD = common.MAV_CMD_REQUEST_PROTOCOL_VERSION
	// Request autopilot capabilities. The receiver should ACK the command and then emit its capabilities in an AUTOPILOT_VERSION message
	MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES MAV_CMD = common.MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES
	// Request camera information (CAMERA_INFORMATION).
	MAV_CMD_REQUEST_CAMERA_INFORMATION MAV_CMD = common.MAV_CMD_REQUEST_CAMERA_INFORMATION
	// Request camera settings (CAMERA_SETTINGS).
	MAV_CMD_REQUEST_CAMERA_SETTINGS MAV_CMD = common.MAV_CMD_REQUEST_CAMERA_SETTINGS
	// Request storage information (STORAGE_INFORMATION). Use the command's target_component to target a specific component's storage.
	MAV_CMD_REQUEST_STORAGE_INFORMATION MAV_CMD = common.MAV_CMD_REQUEST_STORAGE_INFORMATION
	// Format a storage medium. Once format is complete, a STORAGE_INFORMATION message is sent. Use the command's target_component to target a specific component's storage.
	MAV_CMD_STORAGE_FORMAT MAV_CMD = common.MAV_CMD_STORAGE_FORMAT
	// Request camera capture status (CAMERA_CAPTURE_STATUS)
	MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS MAV_CMD = common.MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS
	// Request flight information (FLIGHT_INFORMATION)
	MAV_CMD_REQUEST_FLIGHT_INFORMATION MAV_CMD = common.MAV_CMD_REQUEST_FLIGHT_INFORMATION
	// Reset all camera settings to Factory Default
	MAV_CMD_RESET_CAMERA_SETTINGS MAV_CMD = common.MAV_CMD_RESET_CAMERA_SETTINGS
	// Set camera running mode. Use NaN for reserved values. GCS will send a MAV_CMD_REQUEST_VIDEO_STREAM_STATUS command after a mode change if the camera supports video streaming.
	MAV_CMD_SET_CAMERA_MODE MAV_CMD = common.MAV_CMD_SET_CAMERA_MODE
	// Set camera zoom. Camera must respond with a CAMERA_SETTINGS message (on success).
	MAV_CMD_SET_CAMERA_ZOOM MAV_CMD = common.MAV_CMD_SET_CAMERA_ZOOM
	// Set camera focus. Camera must respond with a CAMERA_SETTINGS message (on success).
	MAV_CMD_SET_CAMERA_FOCUS MAV_CMD = common.MAV_CMD_SET_CAMERA_FOCUS
	// Set that a particular storage is the preferred location for saving photos, videos, and/or other media (e.g. to set that an SD card is used for storing videos).
	// There can only be one preferred save location for each particular media type: setting a media usage flag will clear/reset that same flag if set on any other storage.
	// If no flag is set the system should use its default storage.
	// A target system can choose to always use default storage, in which case it should ACK the command with MAV_RESULT_UNSUPPORTED.
	// A target system can choose to not allow a particular storage to be set as preferred storage, in which case it should ACK the command with MAV_RESULT_DENIED.
	MAV_CMD_SET_STORAGE_USAGE MAV_CMD = common.MAV_CMD_SET_STORAGE_USAGE
	// Tagged jump target. Can be jumped to with MAV_CMD_DO_JUMP_TAG.
	MAV_CMD_JUMP_TAG MAV_CMD = common.MAV_CMD_JUMP_TAG
	// Jump to the matching tag in the mission list. Repeat this action for the specified number of times. A mission should contain a single matching tag for each jump. If this is not the case then a jump to a missing tag should complete the mission, and a jump where there are multiple matching tags should always select the one with the lowest mission sequence number.
	MAV_CMD_DO_JUMP_TAG MAV_CMD = common.MAV_CMD_DO_JUMP_TAG
	// Set gimbal manager pitch/yaw setpoints (low rate command). It is possible to set combinations of the values below. E.g. an angle as well as a desired angular rate can be used to get to this angle at a certain angular rate, or an angular rate only will result in continuous turning. NaN is to be used to signal unset. Note: only the gimbal manager will react to this command - it will be ignored by a gimbal device. Use GIMBAL_MANAGER_SET_PITCHYAW if you need to stream pitch/yaw setpoints at higher rate.
	MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW MAV_CMD = common.MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW
	// Gimbal configuration to set which sysid/compid is in primary and secondary control.
	MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE MAV_CMD = common.MAV_CMD_DO_GIMBAL_MANAGER_CONFIGURE
	// Start image capture sequence. CAMERA_IMAGE_CAPTURED must be emitted after each capture.
	// Param1 (id) may be used to specify the target camera: 0: all cameras, 1 to 6: autopilot-connected cameras, 7-255: MAVLink camera component ID.
	// It is needed in order to target specific cameras connected to the autopilot, or specific sensors in a multi-sensor camera (neither of which have a distinct MAVLink component ID).
	// It is also needed to specify the target camera in missions.
	// When used in a mission, an autopilot should execute the MAV_CMD for a specified local camera (param1 = 1-6), or resend it as a command if it is intended for a MAVLink camera (param1 = 7 - 255), setting the command's target_component as the param1 value (and setting param1 in the command to zero).
	// If the param1 is 0 the autopilot should do both.
	// When sent in a command the target MAVLink address is set using target_component.
	// If addressed specifically to an autopilot: param1 should be used in the same way as it is for missions (though command should NACK with MAV_RESULT_DENIED if a specified local camera does not exist).
	// If addressed to a MAVLink camera, param 1 can be used to address all cameras (0), or to separately address 1 to 7 individual sensors. Other values should be NACKed with MAV_RESULT_DENIED.
	// If the command is broadcast (target_component is 0) then param 1 should be set to 0 (any other value should be NACKED with MAV_RESULT_DENIED). An autopilot would trigger any local cameras and forward the command to all channels.
	MAV_CMD_IMAGE_START_CAPTURE MAV_CMD = common.MAV_CMD_IMAGE_START_CAPTURE
	// Stop image capture sequence.
	// Param1 (id) may be used to specify the target camera: 0: all cameras, 1 to 6: autopilot-connected cameras, 7-255: MAVLink camera component ID.
	// It is needed in order to target specific cameras connected to the autopilot, or specific sensors in a multi-sensor camera (neither of which have a distinct MAVLink component ID).
	// It is also needed to specify the target camera in missions.
	// When used in a mission, an autopilot should execute the MAV_CMD for a specified local camera (param1 = 1-6), or resend it as a command if it is intended for a MAVLink camera (param1 = 7 - 255), setting the command's target_component as the param1 value (and setting param1 in the command to zero).
	// If the param1 is 0 the autopilot should do both.
	// When sent in a command the target MAVLink address is set using target_component.
	// If addressed specifically to an autopilot: param1 should be used in the same way as it is for missions (though command should NACK with MAV_RESULT_DENIED if a specified local camera does not exist).
	// If addressed to a MAVLink camera, param1 can be used to address all cameras (0), or to separately address 1 to 7 individual sensors. Other values should be NACKed with MAV_RESULT_DENIED.
	// If the command is broadcast (target_component is 0) then param 1 should be set to 0 (any other value should be NACKED with MAV_RESULT_DENIED). An autopilot would trigger any local cameras and forward the command to all channels.
	MAV_CMD_IMAGE_STOP_CAPTURE MAV_CMD = common.MAV_CMD_IMAGE_STOP_CAPTURE
	// Re-request a CAMERA_IMAGE_CAPTURED message.
	MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE MAV_CMD = common.MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE
	// Enable or disable on-board camera triggering system.
	MAV_CMD_DO_TRIGGER_CONTROL MAV_CMD = common.MAV_CMD_DO_TRIGGER_CONTROL
	// If the camera supports point visual tracking (CAMERA_CAP_FLAGS_HAS_TRACKING_POINT is set), this command allows to initiate the tracking.
	MAV_CMD_CAMERA_TRACK_POINT MAV_CMD = common.MAV_CMD_CAMERA_TRACK_POINT
	// If the camera supports rectangle visual tracking (CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE is set), this command allows to initiate the tracking.
	MAV_CMD_CAMERA_TRACK_RECTANGLE MAV_CMD = common.MAV_CMD_CAMERA_TRACK_RECTANGLE
	// Stops ongoing tracking.
	MAV_CMD_CAMERA_STOP_TRACKING MAV_CMD = common.MAV_CMD_CAMERA_STOP_TRACKING
	// Starts video capture (recording).
	MAV_CMD_VIDEO_START_CAPTURE MAV_CMD = common.MAV_CMD_VIDEO_START_CAPTURE
	// Stop the current video capture (recording).
	MAV_CMD_VIDEO_STOP_CAPTURE MAV_CMD = common.MAV_CMD_VIDEO_STOP_CAPTURE
	// Start video streaming
	MAV_CMD_VIDEO_START_STREAMING MAV_CMD = common.MAV_CMD_VIDEO_START_STREAMING
	// Stop the given video stream
	MAV_CMD_VIDEO_STOP_STREAMING MAV_CMD = common.MAV_CMD_VIDEO_STOP_STREAMING
	// Request video stream information (VIDEO_STREAM_INFORMATION)
	MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION MAV_CMD = common.MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION
	// Request video stream status (VIDEO_STREAM_STATUS)
	MAV_CMD_REQUEST_VIDEO_STREAM_STATUS MAV_CMD = common.MAV_CMD_REQUEST_VIDEO_STREAM_STATUS
	// Request to start streaming logging data over MAVLink (see also LOGGING_DATA message)
	MAV_CMD_LOGGING_START MAV_CMD = common.MAV_CMD_LOGGING_START
	// Request to stop streaming log data over MAVLink
	MAV_CMD_LOGGING_STOP           MAV_CMD = common.MAV_CMD_LOGGING_STOP
	MAV_CMD_AIRFRAME_CONFIGURATION MAV_CMD = common.MAV_CMD_AIRFRAME_CONFIGURATION
	// Request to start/stop transmitting over the high latency telemetry
	MAV_CMD_CONTROL_HIGH_LATENCY MAV_CMD = common.MAV_CMD_CONTROL_HIGH_LATENCY
	// Create a panorama at the current position
	MAV_CMD_PANORAMA_CREATE MAV_CMD = common.MAV_CMD_PANORAMA_CREATE
	// Request VTOL transition
	MAV_CMD_DO_VTOL_TRANSITION MAV_CMD = common.MAV_CMD_DO_VTOL_TRANSITION
	// Request authorization to arm the vehicle to a external entity, the arm authorizer is responsible to request all data that is needs from the vehicle before authorize or deny the request.
	// If approved the COMMAND_ACK message progress field should be set with period of time that this authorization is valid in seconds.
	// If the authorization is denied COMMAND_ACK.result_param2 should be set with one of the reasons in ARM_AUTH_DENIED_REASON.
	MAV_CMD_ARM_AUTHORIZATION_REQUEST MAV_CMD = common.MAV_CMD_ARM_AUTHORIZATION_REQUEST
	// This command sets the submode to standard guided when vehicle is in guided mode. The vehicle holds position and altitude and the user can input the desired velocities along all three axes.
	MAV_CMD_SET_GUIDED_SUBMODE_STANDARD MAV_CMD = common.MAV_CMD_SET_GUIDED_SUBMODE_STANDARD
	// This command sets submode circle when vehicle is in guided mode. Vehicle flies along a circle facing the center of the circle. The user can input the velocity along the circle and change the radius. If no input is given the vehicle will hold position.
	MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE MAV_CMD = common.MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE
	// Delay mission state machine until gate has been reached.
	MAV_CMD_CONDITION_GATE MAV_CMD = common.MAV_CMD_CONDITION_GATE
	// Fence return point (there can only be one such point in a geofence definition). If rally points are supported they should be used instead.
	MAV_CMD_NAV_FENCE_RETURN_POINT MAV_CMD = common.MAV_CMD_NAV_FENCE_RETURN_POINT
	// Fence vertex for an inclusion polygon (the polygon must not be self-intersecting). The vehicle must stay within this area. Minimum of 3 vertices required.
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION MAV_CMD = common.MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION
	// Fence vertex for an exclusion polygon (the polygon must not be self-intersecting). The vehicle must stay outside this area. Minimum of 3 vertices required.
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION MAV_CMD = common.MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION
	// Circular fence area. The vehicle must stay inside this area.
	MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION MAV_CMD = common.MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION
	// Circular fence area. The vehicle must stay outside this area.
	MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION MAV_CMD = common.MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION
	// Rally point. You can have multiple rally points defined.
	MAV_CMD_NAV_RALLY_POINT MAV_CMD = common.MAV_CMD_NAV_RALLY_POINT
	// Commands the vehicle to respond with a sequence of messages UAVCAN_NODE_INFO, one message per every UAVCAN node that is online. Note that some of the response messages can be lost, which the receiver can detect easily by checking whether every received UAVCAN_NODE_STATUS has a matching message UAVCAN_NODE_INFO received earlier; if not, this command should be sent again in order to request re-transmission of the node information messages.
	MAV_CMD_UAVCAN_GET_NODE_INFO MAV_CMD = common.MAV_CMD_UAVCAN_GET_NODE_INFO
	// Trigger the start of an ADSB-out IDENT. This should only be used when requested to do so by an Air Traffic Controller in controlled airspace. This starts the IDENT which is then typically held for 18 seconds by the hardware per the Mode A, C, and S transponder spec.
	MAV_CMD_DO_ADSB_OUT_IDENT MAV_CMD = common.MAV_CMD_DO_ADSB_OUT_IDENT
	// Deploy payload on a Lat / Lon / Alt position. This includes the navigation to reach the required release position and velocity.
	MAV_CMD_PAYLOAD_PREPARE_DEPLOY MAV_CMD = common.MAV_CMD_PAYLOAD_PREPARE_DEPLOY
	// Control the payload deployment.
	MAV_CMD_PAYLOAD_CONTROL_DEPLOY MAV_CMD = common.MAV_CMD_PAYLOAD_CONTROL_DEPLOY
	// Magnetometer calibration based on provided known yaw. This allows for fast calibration using WMM field tables in the vehicle, given only the known yaw of the vehicle. If Latitude and longitude are both zero then use the current vehicle location.
	MAV_CMD_FIXED_MAG_CAL_YAW MAV_CMD = common.MAV_CMD_FIXED_MAG_CAL_YAW
	// Command to operate winch.
	MAV_CMD_DO_WINCH MAV_CMD = common.MAV_CMD_DO_WINCH
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_1 MAV_CMD = common.MAV_CMD_WAYPOINT_USER_1
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_2 MAV_CMD = common.MAV_CMD_WAYPOINT_USER_2
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_3 MAV_CMD = common.MAV_CMD_WAYPOINT_USER_3
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_4 MAV_CMD = common.MAV_CMD_WAYPOINT_USER_4
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_5 MAV_CMD = common.MAV_CMD_WAYPOINT_USER_5
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_1 MAV_CMD = common.MAV_CMD_SPATIAL_USER_1
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_2 MAV_CMD = common.MAV_CMD_SPATIAL_USER_2
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_3 MAV_CMD = common.MAV_CMD_SPATIAL_USER_3
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_4 MAV_CMD = common.MAV_CMD_SPATIAL_USER_4
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_5 MAV_CMD = common.MAV_CMD_SPATIAL_USER_5
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_1 MAV_CMD = common.MAV_CMD_USER_1
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_2 MAV_CMD = common.MAV_CMD_USER_2
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_3 MAV_CMD = common.MAV_CMD_USER_3
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_4 MAV_CMD = common.MAV_CMD_USER_4
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_5 MAV_CMD = common.MAV_CMD_USER_5
	// Request forwarding of CAN packets from the given CAN bus to this component. CAN Frames are sent using CAN_FRAME and CANFD_FRAME messages
	MAV_CMD_CAN_FORWARD MAV_CMD = common.MAV_CMD_CAN_FORWARD
)
