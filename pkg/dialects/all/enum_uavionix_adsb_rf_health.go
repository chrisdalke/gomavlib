//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/chrisdalke/gomavlib/v3/pkg/dialects/uavionix"
)

// Status flags for ADS-B transponder dynamic output
type UAVIONIX_ADSB_RF_HEALTH = uavionix.UAVIONIX_ADSB_RF_HEALTH

const (
	UAVIONIX_ADSB_RF_HEALTH_INITIALIZING UAVIONIX_ADSB_RF_HEALTH = uavionix.UAVIONIX_ADSB_RF_HEALTH_INITIALIZING
	UAVIONIX_ADSB_RF_HEALTH_OK           UAVIONIX_ADSB_RF_HEALTH = uavionix.UAVIONIX_ADSB_RF_HEALTH_OK
	UAVIONIX_ADSB_RF_HEALTH_FAIL_TX      UAVIONIX_ADSB_RF_HEALTH = uavionix.UAVIONIX_ADSB_RF_HEALTH_FAIL_TX
	UAVIONIX_ADSB_RF_HEALTH_FAIL_RX      UAVIONIX_ADSB_RF_HEALTH = uavionix.UAVIONIX_ADSB_RF_HEALTH_FAIL_RX
)
