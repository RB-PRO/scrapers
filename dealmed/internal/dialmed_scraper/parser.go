package dialmed_scraper

import (
	"github.com/charmbracelet/log"
	"time"
)

const URL = "https://www.deal-med.ru"

func Start() {
	products, err := parse()
	if err != nil {
		panic(err)
	}

	err = saveXlsx(products)
	if err != nil {
		panic(err)
	}
}

func parse() ([]product, error) {
	categories1, err := parseCategoriesOne()
	if err != nil {
		return nil, err
	}

	var categories2 []categoryTwo
	for i, category1 := range categories1 {
		log.Infof("1: [%d/%d] %s\n", i+1, len(categories1), URL+category1.link)
		c2, err := parseCategoriesTwo(category1)
		if err != nil {
			return nil, err
		}
		categories2 = append(categories2, c2...)
		time.Sleep(time.Millisecond * 200)
	}

	var categories3 []categoryThree
	for i, category2 := range categories2 {
		log.Infof("2: [%d/%d] %s\n", i+1, len(categories2), URL+category2.link)
		c3, err := parseCategoriesThree(category2)
		if err != nil {
			return nil, err
		}
		categories3 = append(categories3, c3...)
		time.Sleep(time.Millisecond * 200)
	}

	var products []product
	for i, category3 := range categories3 {
		log.Infof("3: [%d/%d] %s\n", i+1, len(categories3), URL+category3.link)
		prod, err := parseCategoriesProduct(category3)
		if err != nil {
			log.Error(err)
		}
		products = append(products, prod)
		time.Sleep(time.Millisecond * 200)
	}

	return products, nil
}
