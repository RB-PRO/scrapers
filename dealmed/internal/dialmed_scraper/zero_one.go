package dialmed_scraper

import (
	"github.com/gocolly/colly/v2"
)

type categoryZero struct {
	name      string
	link      string
	imageLink string
}

type categoryOne struct {
	categoryZero
	name string
	link string
}

// https://www.deal-med.ru/
func parseCategoriesOne() ([]categoryOne, error) {
	c := colly.NewCollector()

	var categories []categoryOne
	c.OnHTML("#content > div > div > div > div > div > a", func(e *colly.HTMLElement) {
		parentElement := e.DOM.Parent().Parent().Parent().Parent().Find("a:nth-child(1)")
		parentHref, _ := parentElement.Attr("href")
		imageHref, _ := parentElement.Find("div > img").Attr("src")
		parentName := parentElement.Find("div > span > b").Text()
		category0 := categoryZero{
			name:      standardizeSpaces(parentName),
			link:      parentHref,
			imageLink: imageHref,
		}

		category1Name := e.DOM.Text()
		category1Href, _ := e.DOM.Attr("href")
		category1 := categoryOne{
			categoryZero: category0,
			name:         standardizeSpaces(category1Name),
			link:         category1Href,
		}

		categories = append(categories, category1)
	})

	if err := c.Visit(URL); err != nil {
		return nil, err
	}
	return categories, nil
}
