//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package standard

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/minimal"
)

// Micro air vehicle / autopilot classes. This identifies the individual model.
type MAV_AUTOPILOT = minimal.MAV_AUTOPILOT

const (
	// Generic autopilot, full support for everything
	MAV_AUTOPILOT_GENERIC MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_GENERIC
	// Reserved for future use.
	MAV_AUTOPILOT_RESERVED MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_RESERVED
	// SLUGS autopilot, http://slugsuav.soe.ucsc.edu
	MAV_AUTOPILOT_SLUGS MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_SLUGS
	// ArduPilot - Plane/Copter/Rover/Sub/Tracker, https://ardupilot.org
	MAV_AUTOPILOT_ARDUPILOTMEGA MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_ARDUPILOTMEGA
	// OpenPilot, http://openpilot.org
	MAV_AUTOPILOT_OPENPILOT MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_OPENPILOT
	// Generic autopilot only supporting simple waypoints
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY
	// Generic autopilot supporting waypoints and other simple navigation commands
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY
	// Generic autopilot supporting the full mission command set
	MAV_AUTOPILOT_GENERIC_MISSION_FULL MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_GENERIC_MISSION_FULL
	// No valid autopilot, e.g. a GCS or other MAVLink component
	MAV_AUTOPILOT_INVALID MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_INVALID
	// PPZ UAV - http://nongnu.org/paparazzi
	MAV_AUTOPILOT_PPZ MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_PPZ
	// UAV Dev Board
	MAV_AUTOPILOT_UDB MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_UDB
	// FlexiPilot
	MAV_AUTOPILOT_FP MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_FP
	// PX4 Autopilot - http://px4.io/
	MAV_AUTOPILOT_PX4 MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_PX4
	// SMACCMPilot - http://smaccmpilot.org
	MAV_AUTOPILOT_SMACCMPILOT MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_SMACCMPILOT
	// AutoQuad -- http://autoquad.org
	MAV_AUTOPILOT_AUTOQUAD MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_AUTOQUAD
	// Armazila -- http://armazila.com
	MAV_AUTOPILOT_ARMAZILA MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_ARMAZILA
	// Aerob -- http://aerob.ru
	MAV_AUTOPILOT_AEROB MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_AEROB
	// ASLUAV autopilot -- http://www.asl.ethz.ch
	MAV_AUTOPILOT_ASLUAV MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_ASLUAV
	// SmartAP Autopilot - http://sky-drones.com
	MAV_AUTOPILOT_SMARTAP MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_SMARTAP
	// AirRails - http://uaventure.com
	MAV_AUTOPILOT_AIRRAILS MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_AIRRAILS
	// Fusion Reflex - https://fusion.engineering
	MAV_AUTOPILOT_REFLEX MAV_AUTOPILOT = minimal.MAV_AUTOPILOT_REFLEX
)
