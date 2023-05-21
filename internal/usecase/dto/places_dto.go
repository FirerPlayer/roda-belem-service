package dto

type FindNearbyPlacesInputDTO struct {
	Lat    float64
	Lng    float64
	Radius float64
}

type FindNearbyPlacesOutputDTO struct {
	ID                    string   `json:"id"`
	Name                  string   `json:"name"`
	FormatedAddress       string   `json:"formatedAddress"`
	Image                 string   `json:"image"`
	Lat                   float64  `json:"lat"`
	Lng                   float64  `json:"lng"`
	Icon                  string   `json:"icon"`
	Types                 []string `json:"types"`
	OpeningPeriod         []string `json:"openingPeriod"`
	Photos                []string `json:"photo"`
	Rating                float64  `json:"rating"`
	AccessibilityFeatures []string `json:"accessibilityFeatures"`
}
