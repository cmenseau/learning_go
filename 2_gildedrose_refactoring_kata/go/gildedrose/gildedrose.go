package gildedrose

import "fmt"

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i Item) String() string {
	return fmt.Sprint(i.Name, ", ", i.SellIn, ", ", i.Quality)
}

const AGED_BRIE string = "Aged Brie"
const BACKSTAGE_PASS string = "Backstage passes to a TAFKAL80ETC concert"
const SULFURAS string = "Sulfuras, Hand of Ragnaros"
const CONJURED string = "Conjured Mana Cake"

func UpdateAllItems(items []*Item) {
	for i := 0; i < len(items); i++ {
		UpdateItem(items[i])
	}
}

func isLegendaryItem(name string) bool {
	return name == SULFURAS
}

func hasNoValueAfterSellDate(name string) bool {
	return name == BACKSTAGE_PASS
}

func beforeSellDate(sellin int) bool {
	return sellin > -1
}

func addToQualityBetweenBounds(quality int, add int) int {
	return max(min(quality+add, 50), 0)
}

func isMoreValuableWithTime(name string) bool {
	return name == AGED_BRIE || name == BACKSTAGE_PASS
}

func getDegradationByDayFor(name string) int {
	switch {
	case isMoreValuableWithTime(name):
		return +1
	case name == CONJURED:
		return -2
	default:
		return -1
	}
}

func updateQuality(item *Item) {

	if isLegendaryItem(item.Name) {
		return
	}

	var qualityDelta int

	if beforeSellDate(item.SellIn) {
		qualityDelta = getDegradationByDayFor(item.Name)

		if item.Name == BACKSTAGE_PASS {
			if item.SellIn < 5 {
				qualityDelta += 2
			} else if item.SellIn < 10 {
				qualityDelta += 1
			}
		}
	} else {
		qualityDelta = 2 * getDegradationByDayFor(item.Name)

		if hasNoValueAfterSellDate(item.Name) {
			qualityDelta = -item.Quality
		}
	}

	item.Quality = addToQualityBetweenBounds(item.Quality, qualityDelta)
}

func UpdateItem(item *Item) {

	updateSellInDate(item)

	updateQuality(item)
}

func updateSellInDate(item *Item) {
	if isLegendaryItem(item.Name) {
		return
	}

	item.SellIn = item.SellIn - 1
}
