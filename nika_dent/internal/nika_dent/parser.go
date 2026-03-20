package nika_dent

const URL = "https://nika-dent.ru"

func Start() {
	products, err := parse()
	if err != nil {
		panic(err)
	}

	if err = saveXlsx(products); err != nil {
		panic(err)
	}
}

func parse() ([]product, error) {
	categories, err := parseCategoriesOne()
	if err != nil {
		return nil, err
	}

	// categories = categories[1:2]
	// fmt.Println("$+$", categories)

	productsMap, err := parseProducts(categories)
	if err != nil {
		return nil, err
	}

	err = parseItems(productsMap)
	if err != nil {
		return nil, err
	}

	count := 0
	productsMap.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	products := make([]product, 0, count)
	productsMap.Range(func(_, value interface{}) bool {
		products = append(products, value.(product))
		return true
	})
	return products, nil
}
