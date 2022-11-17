package models

type Product struct {
	Name       string `json:"name"`
	Price      string `json:"price"`
	ImageUrl   string `json:"imageUrl"`
	Vendor     string `json:"vendor"`
	ProductUrl string `json:"productUrl"`
}
