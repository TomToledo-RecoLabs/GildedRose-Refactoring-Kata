package main

import (
	"fmt"
)

const (
	agedBrieName  = "Aged Brie"
	backstageName = "Backstage passes to a TAFKAL80ETC concert"
	sulfurasName  = "Sulfuras, Hand of Ragnaros"
	conjuredName  = "Conjured"
)

type Item struct {
	name            string
	sellIn, quality int
}

func (item *Item) Name() string {
	return item.name
}

func (item *Item) increaseQuality(val int) error {
	if val < 0 {
		return fmt.Errorf("increaseQuality: val is negative, %v", val)
	}

	item.quality = item.quality + val
	if item.quality > 50 {
		item.quality = 50
	}
	return nil
}

func (item *Item) decreaseQuality(val int) error {
	if val < 0 {
		return fmt.Errorf("decreaseQuality: val is negative, %v", val)
	}

	item.quality = item.quality - val
	if item.quality < 0 {
		item.quality = 0
	}
	return nil
}

func (item *Item) Quality() int {
	return item.quality
}

func (item *Item) decreaseSellIn(val int) error {
	if val < 0 {
		return fmt.Errorf("decreaseSellIn: val is negative, %v", val)
	}

	item.sellIn = item.sellIn - val
	return nil
}

func (item *Item) SellIn() int {
	return item.sellIn
}

func (item *Item) defaultCase() error {
	item.decreaseSellIn(1)

	if item.SellIn() < 0 {
		item.decreaseQuality(2)
		return nil
	}

	item.decreaseQuality(1)
	return nil
}

func (item *Item) sulfurasCase() error {
	return nil
}

func (item *Item) backstageCase() error {
	item.decreaseSellIn(1)
	if item.SellIn() < 0 {
		item.decreaseQuality(item.Quality())
		return nil
	}

	switch {

	case item.SellIn() < 5:
		item.increaseQuality(3)

	case item.SellIn() < 10:
		item.increaseQuality(2)

	default:
		item.increaseQuality(1)

	}
	return nil
}

func (item *Item) agedBrieCase() error {
	item.decreaseSellIn(1)

	if item.SellIn() < 0 {
		item.increaseQuality(2)
		return nil
	}

	item.increaseQuality(1)
	return nil
}

func (item *Item) conjuredCase() error {
	item.decreaseSellIn(1)

	if item.SellIn() < 0 {
		item.decreaseQuality(4)
		return nil
	}

	item.decreaseQuality(2)
	return nil
}

func (item *Item) update() error {

	switch item.Name() {

	case agedBrieName:
		fmt.Println("start handling", agedBrieName, "Case")
		return item.agedBrieCase()

	case backstageName:
		fmt.Println("start handling", backstageName, "Case")
		return item.backstageCase()

	case sulfurasName:
		fmt.Println("start handling", sulfurasName, "Case")
		return item.sulfurasCase()

	case conjuredName:
		fmt.Println("start handling", conjuredName, "Case")
		return item.conjuredCase()

	default:
		fmt.Println("start handling default Case")
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
