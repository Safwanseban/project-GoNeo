package types

import (
	"time"
)

type OfferCompany struct {
	OfferID            uint      `json:"offerId" gorm:"primaryKey"`
	ClientID           uint      `json:"clientId"`
	Country            string    `json:"country"`
	Image              string    `json:"image"`
	ImageWidth         uint      `json:"imageWidth"`
	ImageHeight        uint      `json:"imageHeight"`
	TextLocale         string    `json:"textLocale"`
	ValidityTextLocale string    `json:"validityTextLocale"`
	Position           int       `json:"position"`
	ValidFrom          time.Time `json:"validFrom"`
	ShowFrom           time.Time `json:"showFrom"`
	ValidTo            time.Time `json:"validTo"`
	Flag               uint      `json:"flag"`
	PageCount          uint      `json:"pageCount"`
	StoreURL           string    `json:"storeUrl"`
	StoreURLTitle      string    `json:"storeUrlTitle"`
	OfferHome          int       `json:"offerHome"`
}
