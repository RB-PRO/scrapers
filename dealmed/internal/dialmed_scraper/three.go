package dialmed_scraper

import "github.com/gocolly/colly/v2"

type categoryThree struct {
	categoryTwo
	parentName string
	name       string
	link       string
}

// https://www.deal-med.ru/rabochee_mesto_oftalmologa.html
func parseCategoriesThree(category2 categoryTwo) ([]categoryThree, error) {
	c := colly.NewCollector()

	var categories []categoryThree
	c.OnHTML("#content > table[class=nobrd] > tbody > tr > td > a", func(e *colly.HTMLElement) {
		category3ParentName := e.DOM.Parent().Parent().Parent().Parent().Prev().Prev().Prev().Text()
		category3Href, _ := e.DOM.Attr("href")
		if len(category3Href) > 0 && category3Href[0] != '/' {
			category3Href = "/" + category3Href
		}
		category3Name := e.DOM.Text()
		if category3Name == "" {
			return
		}

		category3 := categoryThree{
			categoryTwo: category2,
			parentName:  standardizeSpaces(category3ParentName),
			name:        standardizeSpaces(category3Name),
			link:        category3Href,
		}

		categories = append(categories, category3)

	})

	if err := c.Visit(URL + category2.link); err != nil {
		return nil, err
	}
	return categories, nil
}
