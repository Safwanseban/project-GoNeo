package types

import "gorm.io/gorm"

type SampleUser struct {
	gorm.Model

	IssuingCountry string
}
type Product struct {
	gorm.Model
	Name            string  `json:"name"`
	Price           uint    `json:"price"`
	Description     string  `json:"description"`
	Type            string  `json:"type"`
	SpecificCountry Country `json:"specificCountry"`
}
type (
	Country string
	Type    string
)

var (
	US        Country = "us"
	Canada    Country = "canada"
	Bahrain   Country = "bahrain"
	India     Country = "india"
	Vegetable Type    = "vegetable"
	Fruit     Type    = "fruit"
)
