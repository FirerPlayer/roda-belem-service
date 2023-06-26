package entity

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID               uuid.UUID
	GooglePlaceId    string
	Name             string
	FormattedAddress string
	Lat              float64
	Lng              float64
	Icon             string
	Types            []string
	OpeningPeriods   []string
	Photos           []Photo
	Rating           float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewPlace(googlePlaceId string, name string, formatedAddress string, lat float64, lng float64, icon string, types []string, openingPeriods []string) *Place {
	return &Place{
		ID:               uuid.New(),
		GooglePlaceId:    googlePlaceId,
		Name:             name,
		FormattedAddress: formatedAddress,
		Lat:              lat,
		Lng:              lng,
		Icon:             icon,
		Types:            types,
		OpeningPeriods:   openingPeriods,
		Photos:           []Photo{},
		Rating:           0.0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (p *Place) GetRating() float64 {
	return p.Rating
}

// var strToAccessibilityFeaturesEnumMap = map[string]AccessibilityFeaturesEnum{
// 	"Ramp":                      Ramp,
// 	"Elevator":                  Elevator,
// 	"AdaptedBathroom":           AdaptedBathroom,
// 	"BrailleSignage":            BrailleSignage,
// 	"WideCirculationAreas":      WideCirculationAreas,
// 	"ReservedParking":           ReservedParking,
// 	"TactilePaving":             TactilePaving,
// 	"AdaptedTelephones":         AdaptedTelephones,
// 	"VideoIntercom":             VideoIntercom,
// 	"AdaptedTablesCounters":     AdaptedTablesCounters,
// 	"WheelchairAvailability":    WheelchairAvailability,
// 	"SignLanguageCommunication": SignLanguageCommunication,
// 	"GuideDogAllowed":           GuideDogAllowed,
// 	"OnlineAccessibility":       OnlineAccessibility,
// 	"AssistiveTechnologyAccess": AssistiveTechnologyAccess,
// }

// func StrToAccessibilityFeaturesEnum(str string) AccessibilityFeaturesEnum {
// 	return strToAccessibilityFeaturesEnumMap[str]
// }
