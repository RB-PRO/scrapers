package dialmed_scraper

import "strings"

func standardizeSpaces(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "\n", " ")
	str = strings.ReplaceAll(str, "\t", " ")
	str = strings.Join(strings.Fields(str), " ")
	return str
}
