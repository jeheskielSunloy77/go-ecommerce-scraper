package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Name       string `json:"name"`
	Price      string `json:"price"`
	ImageUrl   string `json:"imageUrl"`
	Vendor     string `json:"vendor"`
	ProductUrl string `json:"productUrl"`
}

func Json(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("q")
	splitedQuery := strings.Split(query, " ")
	searchQuery := strings.Join(splitedQuery, "+")

	c := colly.NewCollector()
	url := "https://www.amazon.com/s?k=" + searchQuery
	products := []Product{}

	c.OnHTML("div[data-component-type=s-search-result]", func(e *colly.HTMLElement) {
		product := Product{
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

	json.NewEncoder(w).Encode(products)
}

// func scraper(searchQuery string) []Product {
// 	amazonProducts := amazonScraper(searchQuery)
// 	ebayProducts := ebayScraper(searchQuery)
// 	products := append(amazonProducts, ebayProducts...)
// 	return products
// }

// func amazonScraper(searchQuery string) []Product {
// 	c := colly.NewCollector()
// 	url := "https://www.amazon.com/s?k=" + searchQuery
// 	products := []Product{}

// 	c.OnHTML("div[data-component-type=s-search-result]", func(e *colly.HTMLElement) {
// 		product := Product{
// 			Name:       e.ChildText("span.a-size-medium.a-color-base.a-text-normal"),
// 			Price:      e.ChildText("span.a-price span.a-offscreen"),
// 			ImageUrl:   e.ChildAttr("div.a-section.aok-relative.s-image-fixed-height img", "src"),
// 			ProductUrl: e.ChildAttr("a.a-link-normal.s-underline-text.s-underline-link-text.s-link-style.a-text-normal", "href"),
// 			Vendor:     "Amazon",
// 		}

// 		products = append(products, product)
// 	})

// 	// c.OnHTML("a.s-pagination-item.s-pagination-next.s-pagination-button.s-pagination-separator", func(e *colly.HTMLElement) {
// 	// 	nextPage := e.Request.AbsoluteURL(e.Attr("href"))
// 	// 	c.Visit(nextPage)
// 	// })

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println(r.URL.String())
// 	})

// 	c.OnError(func(_ *colly.Response, err error) {
// 		fmt.Println("Something went wrong:", err)
// 	})

// 	if err := c.Visit(url); err != nil {
// 		fmt.Println("Error on start of crawl: ", err)
// 	}

// 	c.Wait()

// 	return products
// }

// func ebayScraper(searchQuery string) []Product {
// 	c := colly.NewCollector()
// 	url := "https://www.ebay.com/sch/i.html?LH_ItemCondition=1000&_nkw=" + searchQuery
// 	products := []Product{}

// 	c.OnHTML("ul.srp-results.srp-list.clearfix li.s-item.s-item__pl-on-bottom", func(e *colly.HTMLElement) {
// 		product := Product{
// 			Name:       e.ChildText("div.s-item__title span"),
// 			Price:      e.ChildText("span.s-item__price"),
// 			ImageUrl:   e.ChildAttr("div.s-item__image-wrapper img", "src"),
// 			ProductUrl: e.ChildAttr("a.s-item__link", "href"),
// 			Vendor:     "Ebay",
// 		}

// 		products = append(products, product)
// 	})
// 	// c.OnHTML("a.pagination__next.icon-link", func(e *colly.HTMLElement) {
// 	// 	nextPage := e.Request.AbsoluteURL(e.Attr("href"))
// 	// 	c.Visit(nextPage)
// 	// })
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println(r.URL.String())
// 	})
// 	c.OnError(func(_ *colly.Response, err error) {
// 		fmt.Println("Something went wrong:", err)
// 	})

// 	if err := c.Visit(url); err != nil {
// 		fmt.Println("Error on start of crawl: ", err)
// 	}

// 	c.Wait()

// 	return products
// }
