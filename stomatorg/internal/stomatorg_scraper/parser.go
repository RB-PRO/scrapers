package stomatorg_scraper

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
)

func init() {
	if err := playwright.Install(); err != nil {
		panic(err)
	}
}

func parse() ([]product, error) {

	categories, err := parseCategories()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n\n", categories[0])

	page, err := NewPage()
	if err != nil {
		return nil, err
	}

	products, err := page.parseProductLinks(categories)
	if err != nil {
		return nil, err
	}

	fmt.Println("len", len(products))
	fmt.Println("product", products[0])

	products, err = parseProductDetails(products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
