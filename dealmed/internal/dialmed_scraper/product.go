package dialmed_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
)

type product struct {
	categoryThree
	imageLink   string
	price       int
	sku         string
	description string
}

// https://www.deal-med.ru/rabochee_mesto_oftalmologa_stern_talmo.html
func parseCategoriesProduct(category3 categoryThree) (product, error) {
	c := colly.NewCollector()

	var prod product
	c.OnHTML("#content > div", func(e *colly.HTMLElement) {
		prodImageLink, _ := e.DOM.Find("#example_img").Attr("src")
		priceString := e.DOM.Find("div.order > div > span > span.price-value").Text()
		priceString = standardizeSpaces(priceString)
		priceString = strings.ReplaceAll(priceString, " ", "")
		price, _ := strconv.Atoi(priceString)
		sku := e.DOM.Find("div[class=artik]").Text()
		sku = strings.ReplaceAll(sku, "Артикул товара: ", "")

		var description string
		e.DOM.Find("p,h2,ul>li").Each(func(i int, s *goquery.Selection) {
			description += s.Text() + "\n"
		})

		prod = product{
			categoryThree: category3,
			imageLink:     prodImageLink,
			price:         price,
			sku:           sku,
			description:   strings.TrimSpace(description),
		}
	})

	if err := c.Visit(URL + category3.link); err != nil {
		return product{}, err
	}
	return prod, nil
}
