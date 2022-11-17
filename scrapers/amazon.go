package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/jeheskielSunloy77/go-ecommerce-scraper/models"
)

func amazonScraper(searchQuery string) []models.Product {
	c := colly.NewCollector()
	url := "https://www.amazon.com/s?k=" + searchQuery
	products := []models.Product{}

	c.OnHTML("div[data-component-type=s-search-result]", func(e *colly.HTMLElement) {
		product := models.Product{
			Name:       e.ChildText("span.a-size-medium.a-color-base.a-text-normal"),
			Price:      e.ChildText("span.a-price span.a-offscreen"),
			ImageUrl:   e.ChildAttr("div.a-section.aok-relative.s-image-fixed-height img", "src"),
			ProductUrl: e.ChildAttr("a.a-link-normal.s-underline-text.s-underline-link-text.s-link-style.a-text-normal", "href"),
			Vendor:     "Amazon",
		}

		products = append(products, product)
	})

	// c.OnHTML("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator", func(e *colly.HTMLElement) {
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
