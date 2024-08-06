package models

type Personage struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	EarnPercent float64 `json:"earn_percent"`
	EarnPeriod  float64 `json:"earn_period"`
	LifePeriod  float64 `json:"life_period"`
}
