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
	c.OnHTML("#content > table > tbody > tr > td > div > strong > a", func(e *colly.HTMLElement) {
		category2href, _ := e.DOM.Attr("href")
		category2name := e.DOM.Text()
		if category2name != "" {
			category2 := categoryTwo{
				categoryOne: category1,
				name:        standardizeSpaces(category2name),
				link:        category2href,
			}

			categories = append(categories, category2)
		}
	})
	if err := c.Visit(URL + category1.link); err != nil {
		return nil, err
	}
	return categories, nil
}
