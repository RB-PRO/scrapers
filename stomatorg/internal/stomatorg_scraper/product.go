package stomatorg_scraper

type product struct {
	category
	name              string
	link              string
	SKU               string
	manufacturer      string
	countryProduction string
	description       string
	photo             string
	price             int
}
