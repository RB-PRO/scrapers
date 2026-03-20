package nika_dent

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

type product struct {
	category    Category
	name        string
	link        string
	imageLink   string
	price       int
	sku         string
	description string
	manufacture string
}

func digitsOnly(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, s)
}

func parseItems(productsMap sync.Map) error {
	c := colly.NewCollector()

	c.OnHTML(`div.product-info > div:last-child > p > a`, func(e *colly.HTMLElement) {
		manufacture := e.DOM.Text()
		val, ok := productsMap.Load(e.Request.URL.Path)
		if !ok {
			panic("ti eblan, eto ne tot url for manufacture")
		}

		prod := val.(product)
		prod.manufacture = manufacture
		manufacture = strings.TrimSpace(manufacture)
		productsMap.Store(e.Request.URL.Path, prod)
	})

	c.OnHTML(`div.tabs-content > div.tabs-content__item.active > div`, func(e *colly.HTMLElement) {
		description := e.DOM.Text()
		val, ok := productsMap.Load(e.Request.URL.Path)
		if !ok {
			panic("ti eblan, eto ne tot url")
		}

		prod := val.(product)
		description = strings.TrimSpace(description)
		prod.description = description
		productsMap.Store(e.Request.URL.Path, prod)
	})

	productsMap.Range(func(_, value interface{}) bool {
		if prod, ok := value.(product); ok {
			fmt.Println(">>>", URL+prod.link)
			if err := c.Visit(URL + prod.link); err != nil {
				return false
			}
		}

		return true
	})

	c.Wait()
	return nil
}

func parseProducts(categories []Category) (sync.Map, error) {
	c := colly.NewCollector()

	var productsMap sync.Map
	c.OnHTML(`div.product-item`, func(e *colly.HTMLElement) {
		categoryVal := e.Request.Ctx.GetAny("category")
		if categoryVal == nil {
			log.Println("Категория не найдена в контексте запроса")
			return
		}
		category := categoryVal.(Category)

		link := e.ChildAttr("a.item-link", "href")
		name := e.ChildText("a.item-link")
		sku := e.ChildText("div.product-article")
		imageLink := e.ChildAttr("img.item-image", "src")
		if strings.Contains(imageLink, "noimage.jpeg") {
			imageLink = ""
		}

		priceStr := e.ChildText("span.price")
		priceStr = digitsOnly(priceStr)
		priceInt, _ := strconv.Atoi(priceStr)

		prod := product{
			category:  category,
			name:      name,
			link:      link,
			sku:       sku,
			imageLink: imageLink,
			price:     priceInt * 100,
		}

		productsMap.Store(link, prod)

		// вот тут
		//fmt.Printf("$+$ %+v\n", prod)
	})

	// visit on other pages
	c.OnHTML(`div[class=pagination] > a`, func(e *colly.HTMLElement) {
		// return // debug
		queryNextParam, ok := e.DOM.Attr("data-url")
		if !ok {
			return
		}

		nextUrl := URL + e.Request.URL.Path + queryNextParam
		fmt.Println(">>-", nextUrl)
		err := c.Request("GET", nextUrl, nil, e.Request.Ctx, nil)
		if err != nil {
			log.Printf("Ошибка создания запроса для пагинации: %v", err)
			return
		}
	})

	for _, category := range categories {
		visitLink := URL + category.link + "?ajax=Y&PAGEN_1=1"
		fmt.Println(">--", visitLink)

		ctx := colly.NewContext()
		ctx.Put("category", category)

		err := c.Request("GET", visitLink, nil, ctx, nil)
		if err != nil {
			return sync.Map{}, fmt.Errorf("ошибка создания запроса для %s: %w", visitLink, err)
		}
	}

	c.Wait()

	return productsMap, nil
}
