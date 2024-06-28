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

func UpdateItem(item *Item) {

	updateSellInDate(item)

	// handle products decreasing in quality first
	if item.Name != AGED_BRIE && item.Name != BACKSTAGE_PASS {
		if item.Quality > 0 {
			if item.Name != SULFURAS {
				item.Quality = item.Quality - 1
			}
		}
	} else { // handle special products, increasing in quality
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.Name == BACKSTAGE_PASS {
				if item.SellIn < 10 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
				if item.SellIn < 5 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
			}
		}
	}

	if item.SellIn < 0 {
		if item.Name != AGED_BRIE {
			if item.Name != BACKSTAGE_PASS {
				if item.Quality > 0 {
					if item.Name != SULFURAS {
						// normal case : decrease twice as fast
						item.Quality = item.Quality - 1
					}
				}
				// backstage pass quality is 0 after the concert
			} else {
				item.Quality = 0
			}
		} else { // aged brie increases twice as fat afer sell by date
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
}

func updateSellInDate(item *Item) {
	if item.Name != SULFURAS {
		item.SellIn = item.SellIn - 1
	}
}
