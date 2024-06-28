package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

type SubtestUpdateItem struct {
	name                string
	item                gildedrose.Item
	expectedItemNextDay gildedrose.Item
}

func (s SubtestUpdateItem) RunSubtest() func(t *testing.T) {

	return func(t *testing.T) {

		originalItem := gildedrose.Item{s.item.Name, s.item.SellIn, s.item.Quality}

		gildedrose.UpdateItem(&s.item)

		if s.expectedItemNextDay != s.item {
			t.Errorf("For %s: Expected %s but got %s ", originalItem, s.expectedItemNextDay, s.item)
		}
	}
}

func TestUpdateItem(t *testing.T) {
	var subtests = []SubtestUpdateItem{
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
			item:                gildedrose.Item{gildedrose.AGED_BRIE, 42, 1},
			expectedItemNextDay: gildedrose.Item{gildedrose.AGED_BRIE, 41, 2},
		},
		{
			name:                "Aged Brie quality increases, even after sell date passed, but twice faster",
			item:                gildedrose.Item{gildedrose.AGED_BRIE, -2, 10},
			expectedItemNextDay: gildedrose.Item{gildedrose.AGED_BRIE, -3, 12},
		},
		{
			name:                "Aged Brie quality increases, never more than 50",
			item:                gildedrose.Item{gildedrose.AGED_BRIE, -2, 50},
			expectedItemNextDay: gildedrose.Item{gildedrose.AGED_BRIE, -3, 50},
		},
		{
			name:                "Sulfuras doesn't have to be sold, and has constant quality",
			item:                gildedrose.Item{gildedrose.SULFURAS, 0, 80},
			expectedItemNextDay: gildedrose.Item{gildedrose.SULFURAS, 0, 80},
		},
		{
			name:                "Sulfuras doesn't have to be sold, and has constant quality",
			item:                gildedrose.Item{gildedrose.SULFURAS, -1, 80},
			expectedItemNextDay: gildedrose.Item{gildedrose.SULFURAS, -1, 80},
		},
		{
			name:                "Backstage passes quality is 0 after concert",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 0, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, -1, 0},
		},
		{
			name:                "Backstage passes quality increases by 2 when there's between 6 and 10 days left, upper limit value",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 10, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 9, 17},
		},
		{
			name:                "Backstage passes quality increases by 2 when there's between 6 and 10 days left",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 7, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 6, 17},
		},
		{
			name:                "Backstage passes quality increases by 2 when there's between 6 and 10 days left, lower limit value",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 6, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 5, 17},
		},
		{
			name:                "Backstage passes quality increases by 3 when there's between 0 and 5 days left, lower limit value",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 1, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 0, 18},
		},
		{
			name:                "Backstage passes quality increases by 3 when there's between 0 and 5 days left, upper limit value",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 4, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 3, 18},
		},
		{
			name:                "Backstage passes quality increases by 3 when there's between 0 and 5 days left - max 50",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 4, 48},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 3, 50},
		},
		{
			name:                "Backstage passes with more than 10 days left : increase in quality, lower limit value",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 11, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 10, 16},
		},
		{
			name:                "Backstage passes with more than 10 days left : increase in quality",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 14, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 13, 16},
		},
		{
			name:                "Backstage passes with more than 10 days left : increase in quality - max 50",
			item:                gildedrose.Item{gildedrose.BACKSTAGE_PASS, 14, 50},
			expectedItemNextDay: gildedrose.Item{gildedrose.BACKSTAGE_PASS, 13, 50},
		},
		{
			name:                "Conjured degrades in quality by 2 each day",
			item:                gildedrose.Item{gildedrose.CONJURED, 14, 5},
			expectedItemNextDay: gildedrose.Item{gildedrose.CONJURED, 13, 3},
		},
		{
			name:                "Conjured degrades in quality by 2 each day, lower sellin limit for normal degradation",
			item:                gildedrose.Item{gildedrose.CONJURED, 1, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.CONJURED, 0, 13},
		},
		{
			name:                "Conjured degrades in quality by 2 each day, uper sellin limit for twice faster degradation",
			item:                gildedrose.Item{gildedrose.CONJURED, 0, 15},
			expectedItemNextDay: gildedrose.Item{gildedrose.CONJURED, -1, 11},
		},
		{
			name:                "Conjured degrades in quality by 2 each day, twice as fast after sellin passed",
			item:                gildedrose.Item{gildedrose.CONJURED, -4, 18},
			expectedItemNextDay: gildedrose.Item{gildedrose.CONJURED, -5, 14},
		},
	}

	for _, st := range subtests {
		t.Run(st.name, st.RunSubtest())
	}
}

func TestUpdateQuality(t *testing.T) {

	items := []*gildedrose.Item{
		{gildedrose.CONJURED, 0, 15},
		{gildedrose.AGED_BRIE, 42, 1},
		{"+5 Dexterity Vest", 10, 20},
		{gildedrose.SULFURAS, 0, 80},
		{gildedrose.SULFURAS, -1, 80},
	}

	itemsCopy := make([]*gildedrose.Item, 0, len(items))
	for _, item := range items {
		itemsCopy = append(itemsCopy, &gildedrose.Item{item.Name, item.SellIn, item.Quality})
	}

	gildedrose.UpdateAllItems(items)

	for idx := range items {
		if gildedrose.UpdateItem(itemsCopy[idx]); *itemsCopy[idx] != *items[idx] {
			t.Errorf("Expected same value but got %s and %s ", itemsCopy[idx], items[idx])
		}
	}
}
