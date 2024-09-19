package dialmed_scraper

import (
	"github.com/gocolly/colly/v2"
)

type categoryTwo struct {
	categoryOne
	name string
	link string
}

// https://www.deal-med.ru/standart_osnashenia_kabineta_oftalmologa.html
func parseCategoriesTwo(category1 categoryOne) ([]categoryTwo, error) {
	c := colly.NewCollector()

	var categories []categoryTwo
	c.OnHTML("table > tbody > tr > td > div > strong > a", func(e *colly.HTMLElement) {
		category2Href, _ := e.DOM.Attr("href")
		if len(category2Href) > 0 && category2Href[0] != '/' {
			category2Href = "/" + category2Href
		}
		category2name := e.DOM.Text()
		if category2name == "" {
			return
		}

		category2 := categoryTwo{
			categoryOne: category1,
			name:        standardizeSpaces(category2name),
			link:        category2Href,
		}

		categories = append(categories, category2)
	})
	if err := c.Visit(URL + category1.link); err != nil {
		return nil, err
	}
	return categories, nil
}
