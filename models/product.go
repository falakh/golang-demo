package models

type Product struct {
	ID              uint `json:"id" gorm:"primary_key"`
	Quantity        int  `json:"quantity"`
	Price           int  `json:"price"`
	CompetitorPrice int  `json:"competitorPrice"`
}
