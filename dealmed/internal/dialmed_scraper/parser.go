package dialmed_scraper

import (
	"fmt"
	"github.com/xuri/excelize/v2"
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
		fmt.Printf("[%d/%d]: 1\n", i+1, len(categories1))
		c2, err := parseCategoriesTwo(category1)
		if err != nil {
			return nil, err
		}
		categories2 = append(categories2, c2...)
		time.Sleep(time.Millisecond * 200)
	}

	var categories3 []categoryThree
	for i, category2 := range categories2 {
		fmt.Printf("[%d/%d]: 2\n", i+1, len(categories2))
		c3, err := parseCategoriesThree(category2)
		if err != nil {
			return nil, err
		}
		categories3 = append(categories3, c3...)
		time.Sleep(time.Millisecond * 200)
	}

	var products []product
	for i, category3 := range categories3 {
		fmt.Printf("[%d/%d]: 3\n", i+1, len(categories3))
		prod, err := parseCategoriesProduct(category3)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
		time.Sleep(time.Millisecond * 200)
	}

	return products, nil
}

func saveXlsx(products []product) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	sheetName := "main"
	if _, err := f.NewSheet(sheetName); err != nil {
		return err
	}
	if err := f.DeleteSheet("Sheet1"); err != nil {
		return err
	}

	// write content
	_ = f.SetCellValue(sheetName, fmt.Sprintf("A%d", 1), "Категория")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("B%d", 1), "Ссылка")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("C%d", 1), "Ссылка на картинку")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("D%d", 1), "ПодКатегория")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("E%d", 1), "Ссылка")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("F%d", 1), "Ссылка на картинку")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("G%d", 1), "Название типа прибора")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("H%d", 1), "Ссылка")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("I%d", 1), "Ссылка на картинку")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("J%d", 1), "Название категории товаров")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("K%d", 1), "Название товара")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("L%d", 1), "Ссылка")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("M%d", 1), "Ссылка на картинку")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("N%d", 1), "Ссылка на картинку")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("O%d", 1), "Цена")

	for i, prod := range products {
		_ = f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), prod.categoryZero.name)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), URL+prod.categoryZero.link)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), URL+prod.categoryZero.imageLink)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), prod.categoryOne.name)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), URL+prod.categoryOne.link)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), URL+prod.categoryOne.imageLink)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), prod.categoryTwo.name)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), URL+prod.categoryTwo.link)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), URL+prod.categoryTwo.imageLink)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), prod.categoryThree.parentName)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), prod.categoryThree.name)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("L%d", i+2), URL+prod.categoryThree.link)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("M%d", i+2), URL+prod.categoryThree.imageLink)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("N%d", i+2), URL+prod.imageLink)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("O%d", i+2), prod.price)
	}

	if err := f.SaveAs("dial_med.xlsx"); err != nil {
		return err
	}
	return nil
}
