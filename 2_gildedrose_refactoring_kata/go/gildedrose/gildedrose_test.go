package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

type Subtest struct {
	name                string
	item                gildedrose.Item
	expectedItemNextDay gildedrose.Item
}

func (s Subtest) RunSubtest() func(t *testing.T) {

	return func(t *testing.T) {

		originalItem := gildedrose.Item{s.item.Name, s.item.SellIn, s.item.Quality}

		gildedrose.UpdateQuality([]*gildedrose.Item{&s.item})

		if s.expectedItemNextDay != s.item {
			t.Errorf("For %s: Expected %s but got %s ", originalItem, s.expectedItemNextDay, s.item)
		}
	}
}

func TestRegularItems(t *testing.T) {
	var subtests = []Subtest{
		{
			name:                "Regular item",
			item:                gildedrose.Item{"Regular", 2, 2},
			expectedItemNextDay: gildedrose.Item{"Regular", 1, 1},
		},
		{
			name:                "Regular item, lower sellin limit for normal degradation",
			item:                gildedrose.Item{"Regular", 1, 2},
			expectedItemNextDay: gildedrose.Item{"Regular", 0, 1},
		},
		{
			name:                "Regular item, uper sellin limit for twice faster degradation",
			item:                gildedrose.Item{"Regular", 0, 2},
			expectedItemNextDay: gildedrose.Item{"Regular", -1, 0},
		},
		{
			name:                "Regular item, sell by date passed -> twice faster degradation : -2",
			item:                gildedrose.Item{"Regular", -1, 10},
			expectedItemNextDay: gildedrose.Item{"Regular", -2, 8},
		},
		{
			name:                "Quality never negative",
			item:                gildedrose.Item{"Regular", 5, 0},
			expectedItemNextDay: gildedrose.Item{"Regular", 4, 0},
		},
		{
			name:                "Quality never negative, even after sell date passed",
			item:                gildedrose.Item{"Regular", -5, 0},
			expectedItemNextDay: gildedrose.Item{"Regular", -6, 0},
		},
		{
			name:                "Aged Brie quality increases",
			item:                gildedrose.Item{"Aged Brie", 42, 1},
			expectedItemNextDay: gildedrose.Item{"Aged Brie", 41, 2},
		},
		{
			name:                "Aged Brie quality increases, even after sell date passed, but twice faster",
			item:                gildedrose.Item{"Aged Brie", -2, 10},
			expectedItemNextDay: gildedrose.Item{"Aged Brie", -3, 12},
		},
		{
			name:                "Aged Brie quality increases, never more than 50",
			item:                gildedrose.Item{"Aged Brie", -2, 50},
			expectedItemNextDay: gildedrose.Item{"Aged Brie", -3, 50},
		},
		{
			name:                "Sulfuras doesn't have to be sold, and has constant quality",
			item:                gildedrose.Item{"Sulfuras, Hand of Ragnaros", 0, 80},
			expectedItemNextDay: gildedrose.Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		},
		{
			name:                "Sulfuras doesn't have to be sold, and has constant quality",
			item:                gildedrose.Item{"Sulfuras, Hand of Ragnaros", -1, 80},
			expectedItemNextDay: gildedrose.Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		},
		{
			name:                "Backstage passes quality is 0 after concert",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 0, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		},
		{
			name:                "Backstage passes quality increases by 2 when there's between 6 and 10 days left, upper limit value",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 10, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 9, 17},
		},
		{
			name:                "Backstage passes quality increases by 2 when there's between 6 and 10 days left",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 7, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 6, 17},
		},
		{
			name:                "Backstage passes quality increases by 2 when there's between 6 and 10 days left, lower limit value",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 6, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 5, 17},
		},
		{
			name:                "Backstage passes quality increases by 3 when there's between 0 and 5 days left, lower limit value",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 1, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 0, 18},
		},
		{
			name:                "Backstage passes quality increases by 3 when there's between 0 and 5 days left, upper limit value",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 4, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 3, 18},
		},
		{
			name:                "Backstage passes quality increases by 3 when there's between 0 and 5 days left - max 50",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 4, 48},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		},
		{
			name:                "Backstage passes with more than 10 days left : increase in quality, lower limit value",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 11, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 10, 16},
		},
		{
			name:                "Backstage passes with more than 10 days left : increase in quality",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 14, 15},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 13, 16},
		},
		{
			name:                "Backstage passes with more than 10 days left : increase in quality - max 50",
			item:                gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 14, 50},
			expectedItemNextDay: gildedrose.Item{"Backstage passes to a TAFKAL80ETC concert", 13, 50},
		},
		{
			name:                "Conjured degrades in quality by 2 each day",
			item:                gildedrose.Item{"Conjured Mana Cake", 14, 5},
			expectedItemNextDay: gildedrose.Item{"Conjured Mana Cake", 13, 3},
		},
		{
			name:                "Conjured degrades in quality by 2 each day, lower sellin limit for normal degradation",
			item:                gildedrose.Item{"Conjured Mana Cake", 1, 15},
			expectedItemNextDay: gildedrose.Item{"Conjured Mana Cake", 0, 13},
		},
		{
			name:                "Conjured degrades in quality by 2 each day, uper sellin limit for twice faster degradation",
			item:                gildedrose.Item{"Conjured Mana Cake", 0, 15},
			expectedItemNextDay: gildedrose.Item{"Conjured Mana Cake", -1, 11},
		},
		{
			name:                "Conjured degrades in quality by 2 each day, twice as fast after sellin passed",
			item:                gildedrose.Item{"Conjured Mana Cake", -4, 18},
			expectedItemNextDay: gildedrose.Item{"Conjured Mana Cake", -5, 14},
		},
	}

	for _, st := range subtests {
		t.Run(st.name, st.RunSubtest())
	}
}
