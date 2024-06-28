package gildedrose

import "fmt"

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i Item) String() string {
	return fmt.Sprint(i.Name, " ", i.SellIn, i.Quality)
}

const AGED_BRIE string = "Aged Brie"
const BACKSTAGE_PASS string = "Backstage passes to a TAFKAL80ETC concert"
const SULFURAS string = "Sulfuras, Hand of Ragnaros"
const CONJURED string = "Conjured Mana Cake"

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		UpdateItem(items[i])
	}
}

func isMoreValuableWithTime(name string) bool {
	return name == AGED_BRIE || name == BACKSTAGE_PASS
}

func hasNoValueAfterSellIn(name string) bool {
	return name == BACKSTAGE_PASS
}

func updateQuality(item *Item) {

	if item.Name == SULFURAS {
		return
	}

	addToQuality := -1
	isPastSellIn := item.SellIn < 0

	if isMoreValuableWithTime(item.Name) {
		addToQuality = -addToQuality
	}

	if item.Name == BACKSTAGE_PASS {
		if item.SellIn < 10 && item.Quality < 50 {
			addToQuality += 1
		}
		if item.SellIn < 5 && item.Quality < 50 {
			addToQuality += 1
		}
	}

	if isPastSellIn {
		addToQuality = 2 * addToQuality

		if hasNoValueAfterSellIn(item.Name) {
			addToQuality = -item.Quality
		}
	}

	item.Quality = max(min(item.Quality+addToQuality, 50), 0)
}

func UpdateItem(item *Item) {

	updateSellInDate(item)

	updateQuality(item)
}

func updateSellInDate(item *Item) {
	if item.Name != SULFURAS {
		item.SellIn = item.SellIn - 1
	}
}
