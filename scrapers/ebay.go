package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/jeheskielSunloy77/go-ecommerce-scraper/models"
)

func ebayScraper(searchQuery string) []models.Product {
	c := colly.NewCollector()
	url := "https://www.ebay.com/sch/i.html?LH_ItemCondition=1000&_nkw=" + searchQuery
	products := []models.Product{}

	c.OnHTML("ul.srp-results.srp-list.clearfix li.s-item.s-item__pl-on-bottom", func(e *colly.HTMLElement) {
		product := models.Product{
			Name:       e.ChildText("div.s-item__title span"),
			Price:      e.ChildText("span.s-item__price"),
			ImageUrl:   e.ChildAttr("div.s-item__image-wrapper img", "src"),
			ProductUrl: e.ChildAttr("a.s-item__link", "href"),
			Vendor:     "Ebay",
		}

		products = append(products, product)
	})
	// c.OnHTML("a.pagination__next.icon-link", func(e *colly.HTMLElement) {
	// 	nextPage := e.Request.AbsoluteURL(e.Attr("href"))
	// 	c.Visit(nextPage)
	// })
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	if err := c.Visit(url); err != nil {
		fmt.Println("Error on start of crawl: ", err)
	}

	c.Wait()

	return products
}
