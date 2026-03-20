package nika_dent

import (
	"github.com/gocolly/colly/v2"
)

type BaseCategory struct {
	name string
}

type Category struct {
	baseCategory BaseCategory
	name         string
	link         string
}

func parseCategoriesOne() ([]Category, error) {
	c := colly.NewCollector()

	var categories []Category
	c.OnHTML("nav > ul > li > ul > li > a", func(e *colly.HTMLElement) {
		parentElement := e.DOM.Parent().Parent().Parent().ChildrenFiltered("a")
		baseCategory := BaseCategory{name: parentElement.Text()}
		link, _ := e.DOM.Attr("href")
		category := Category{
			baseCategory: baseCategory,
			name:         e.DOM.Text(),
			link:         link,
		}
		categories = append(categories, category)
	})

	if err := c.Visit(URL + "/catalog/khirurgiya/shovnye-materialy"); err != nil {
		return nil, err
	}
	return categories, nil
}
