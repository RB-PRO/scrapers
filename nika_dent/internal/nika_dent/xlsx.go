package nika_dent

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// saveXlsx сохраняет список товаров в файл nika-dent.xlsx
func saveXlsx(products []product) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			// В примере использовался panic, можно заменить на логирование
			panic(err)
		}
	}()

	// Создаём новый лист с именем "main"
	sheetName := "main"
	idx, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("не удалось создать лист: %w", err)
	}
	// Удаляем стандартный лист "Sheet1" (если он есть)
	_ = f.DeleteSheet("Sheet1")
	f.SetActiveSheet(idx)

	// Заголовки колонок
	headers := []string{
		"Базовая категория",
		"Категория",
		"Ссылка на категорию",
		"Название товара",
		"Ссылка на товар",
		"Ссылка на картинку",
		"Цена",
		"Артикул",
		"Производитель",
		"Описание",
	}
	for col, h := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+col, 1) // A1, B1, C1, ...
		if err := f.SetCellValue(sheetName, cell, h); err != nil {
			return err
		}
	}

	for i, p := range products {
		row := i + 2
		_ = f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), p.category.baseCategory.name)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), p.category.name)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), URL+p.category.link)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), p.name)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), URL+p.link)
		if p.imageLink != "" {
			_ = f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), URL+p.imageLink)
		}
		_ = f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), p.price)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), p.sku)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), p.manufacture)
		_ = f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), p.description)
	}

	if err := f.SaveAs("nika-dent.xlsx"); err != nil {
		return fmt.Errorf("ошибка сохранения файла: %w", err)
	}
	return nil
}
