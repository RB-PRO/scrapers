package dialmed_scraper

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

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
	_ = f.SetCellValue(sheetName, fmt.Sprintf("P%d", 1), "Артикул")
	_ = f.SetCellValue(sheetName, fmt.Sprintf("Q%d", 1), "Описание")

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
		_ = f.SetCellValue(sheetName, fmt.Sprintf("P%d", i+2), prod.sku)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("Q%d", i+2), prod.description)
	}

	if err := f.SaveAs("dial_med.xlsx"); err != nil {
		return err
	}
	return nil
}
