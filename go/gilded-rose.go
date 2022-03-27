package main

import (
	"fmt"
)

type Item struct {
	name            string
	sellIn, quality int
}

func (item *Item) Name() string {
	return item.name
}

func (item *Item) addQuality(val int) error {
	if val < 0 {
		return fmt.Errorf("addQuality: val is negative, %v", val)
	}

	item.quality = item.quality + val
	return nil
}

func (item *Item) subQuality(val int) error {
	if val < 0 {
		return fmt.Errorf("subQuality: val is negative, %v", val)
	}

	item.quality = item.quality - val
	return nil
}

func (item *Item) incQuality() error {
	return item.addQuality(1)
}

func (item *Item) decQuality() error {
	return item.subQuality(1)
}

func (item *Item) resetQuality() error {
	item.quality = 0
	return nil
}

func (item *Item) Quality() int {
	return item.quality
}

func (item *Item) subSellIn(val int) error {
	if val < 0 {
		return fmt.Errorf("subSellIn: val is negative, %v", val)
	}

	item.sellIn = item.sellIn - val
	return nil
}

func (item *Item) decSellIn() error {
	return item.subSellIn(1)
}

func (item *Item) SellIn() int {
	return item.sellIn
}

func (item *Item) defaultCase() error {
	if item.Quality() > 0 {
		item.decQuality()
		if item.SellIn() <= 0 && item.Quality() > 0 {
			item.decQuality()
		}
	}
	item.decSellIn()

	return nil
}

func (item *Item) sulfurasCase() error {
	return nil
}

func (item *Item) backstageCase() error {
	item.decSellIn()
	if item.SellIn() < 0 {
		item.resetQuality()
		return nil
	}

	if item.Quality() < 50 {
		item.incQuality()
		if item.SellIn() < 10 && item.Quality() < 50 {
			item.incQuality()
		}
		if item.SellIn() < 5 && item.Quality() < 50 {
			item.incQuality()
		}
	}

	return nil
}

func (item *Item) agedBrieCase() error {
	if item.Quality() < 50 {
		item.incQuality()
		if item.SellIn() <= 0 && item.Quality() < 50 {
			item.incQuality()
		}
	}
	item.decSellIn()
	return nil
}

func (item *Item) conjuredCase() error {
	item.subQuality(2)
	if item.SellIn() <= 0 {
		item.subQuality(2)
	}
	if item.Quality() < 0 {
		item.resetQuality()
	}
	item.decSellIn()
	return nil
}

func (item *Item) update() error {

	switch item.Name() {

	case "Aged Brie":
		return item.agedBrieCase()

	case "Backstage passes to a TAFKAL80ETC concert":
		return item.backstageCase()

	case "Sulfuras, Hand of Ragnaros":
		return item.sulfurasCase()

	case "Conjured":
		return item.conjuredCase()

	default:
		return item.defaultCase()

	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		if err := items[i].update(); err != nil {
			fmt.Println("item", items[i], ",update Error", err.Error())
		}
	}

}
