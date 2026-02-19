package stomatorg_scraper

import (
	"fmt"
)

const URL = "https://stomatorg.ru"

func Start() {
	products, err := parse()

	if err != nil {
		//panic(err)
		fmt.Println("err:", err)
	}
	_ = products

	//err = saveXlsx(products)
	//if err != nil {
	//	panic(err)
	//}

}
