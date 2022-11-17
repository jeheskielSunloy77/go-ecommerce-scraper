package scrapers

import "github.com/jeheskielSunloy77/go-ecommerce-scraper/models"

func Scraper(searchQuery string) []models.Product {
	amazonProducts := amazonScraper(searchQuery)
	ebayProducts := ebayScraper(searchQuery)
	products := append(amazonProducts, ebayProducts...)
	return products
}
