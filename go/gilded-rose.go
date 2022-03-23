package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

func (item *Item) Name() string {
	return item.name
}

func (item *Item) addQuality(val int) bool {
	item.quality = item.quality + val
	return true
}

func (item *Item) incQuality() bool {
	item.addQuality(1)
	return true
}

func (item *Item) decQuality() bool {
	item.addQuality(-1)
	return true
}

func (item *Item) resetQuality() bool {
	item.quality = 0
	return true
}

func (item *Item) Quality() int {
	return item.quality
}

func (item *Item) addSellIn(val int) bool {
	item.sellIn = item.sellIn + val
	return true
}

func (item *Item) decSellIn() bool {
	item.addSellIn(-1)
	return true
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

func (item *Item) update() error {

	switch item.Name() {

	case "Aged Brie":
		return item.agedBrieCase()

	case "Backstage passes to a TAFKAL80ETC concert":
		return item.backstageCase()

	case "Sulfuras, Hand of Ragnaros":
		return item.sulfurasCase()

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
