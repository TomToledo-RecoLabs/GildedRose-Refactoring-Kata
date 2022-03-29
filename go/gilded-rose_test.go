package main

import (
	"fmt"
	"reflect"
	"testing"
)

type itemFieldChangeTest struct {
	testname string
	in       *Item
	change   int
	want     error
}

func TestUpdateQuality(t *testing.T) {
	var tests = []struct {
		testname string
		in       []*Item
		want     []*Item
	}{
		{
			testname: "Sanity success test: A 1 1",
			in:       []*Item{{name: "A", sellIn: 1, quality: 1}},
			want:     []*Item{{name: "A", sellIn: 0, quality: 0}},
		},
		{
			testname: "Sanity success test: quality is 0",
			in:       []*Item{{name: "A", sellIn: 1, quality: 0}},
			want:     []*Item{{name: "A", sellIn: 0, quality: 0}},
		},
		{
			testname: "Sanity success test: stay the same",
			in:       []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 1}},
			want:     []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 1}},
		},
		{
			testname: "Sanity success test: max quality 50, starting with 50",
			in:       []*Item{{name: "Aged Brie", sellIn: 1, quality: 50}},
			want:     []*Item{{name: "Aged Brie", sellIn: 0, quality: 50}},
		},
		{
			testname: "Sanity success test: max quality 50, starting with 49",
			in:       []*Item{{name: "Aged Brie", sellIn: 1, quality: 49}},
			want:     []*Item{{name: "Aged Brie", sellIn: 0, quality: 50}},
		},
		{
			testname: "Sanity success test: Backstage max quality 50",
			in:       []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 20, quality: 49}},
			want:     []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 19, quality: 50}},
		},
		{
			testname: "Sanity success test: Backstage max quality 50, with sellin 5",
			in:       []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 49}},
			want:     []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50}},
		},
		{
			testname: "Sanity success test: Backstage with 5 selling, increcmenting by 3",
			in:       []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 40}},
			want:     []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 43}},
		},
		{
			testname: "Sanity success test: Backstage with 10 selling, increcmenting by 2",
			in:       []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 40}},
			want:     []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 42}},
		}, //end of first section
		{
			testname: "Sanity success test: stay the same Sulfuras with negative sellin -1",
			in:       []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: -1, quality: 1}},
			want:     []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: -1, quality: 1}},
		},
		{
			testname: "Sanity success test: sellin 0 increamented by 2",
			in:       []*Item{{name: "Aged Brie", sellIn: 0, quality: 40}},
			want:     []*Item{{name: "Aged Brie", sellIn: -1, quality: 42}},
		},
		{
			testname: "Sanity success test: Backstage reset",
			in:       []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 0, quality: 10}},
			want:     []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: -1, quality: 0}},
		},
		{
			testname: "Sanity success test: general decreamenting by 2",
			in:       []*Item{{name: "A", sellIn: 0, quality: 10}},
			want:     []*Item{{name: "A", sellIn: -1, quality: 8}},
		},
		{
			testname: "Sanity success test: Conjured 0 10, decreamenting by 4",
			in:       []*Item{{name: "Conjured", sellIn: 0, quality: 10}},
			want:     []*Item{{name: "Conjured", sellIn: -1, quality: 6}},
		},
		{
			testname: "Sanity success test: Conjured 1 10, decreamenting by 2",
			in:       []*Item{{name: "Conjured", sellIn: 1, quality: 10}},
			want:     []*Item{{name: "Conjured", sellIn: 0, quality: 8}},
		},
		{
			testname: "Sanity success test: Conjured 0 10, decreamenting by 2",
			in:       []*Item{{name: "Conjured", sellIn: 0, quality: 2}},
			want:     []*Item{{name: "Conjured", sellIn: -1, quality: 0}},
		},
	}

	for testInd, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			UpdateQuality(tt.in)
			for i := range tt.in {
				if !reflect.DeepEqual(tt.in[i], tt.want[i]) {
					t.Errorf("test index:%v, got %v, want %v", testInd, tt.in[i], tt.want[i])
				}
			}
		})
	}
}

func TestIncreaseQuality(t *testing.T) {
	var tests = []itemFieldChangeTest{
		{testname: "A 1 1, increaseQuality -1", in: &Item{name: "A", sellIn: 1, quality: 1}, change: -1, want: fmt.Errorf("increaseQuality: val is negative, %v", -1)},
		{testname: "A 1 1, increaseQuality 1", in: &Item{name: "A", sellIn: 1, quality: 1}, change: 1, want: nil},
		{testname: "A 1 50, increaseQuality 1", in: &Item{name: "A", sellIn: 1, quality: 50}, change: 1, want: nil},
	}

	for testInd, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := tt.in.increaseQuality(tt.change)
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("test index:%v, got %v, want %v", testInd, err, tt.want)
			}
		})
	}
}

func TestDecreaseQuality(t *testing.T) {
	var tests = []itemFieldChangeTest{
		{testname: "A 1 1, decreaseQuality -1", in: &Item{name: "A", sellIn: 1, quality: 1}, change: -1, want: fmt.Errorf("decreaseQuality: val is negative, %v", -1)},
		{testname: "A 1 1, decreaseQuality 1", in: &Item{name: "A", sellIn: 1, quality: 1}, change: 1, want: nil},
		{testname: "A 1 0, decreaseQuality 1", in: &Item{name: "A", sellIn: 1, quality: 0}, change: 1, want: nil},
	}

	for testInd, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := tt.in.decreaseQuality(tt.change)
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("test index:%v, got %v, want %v", testInd, err, tt.want)
			}
		})
	}
}

func TestDecreaseSellIn(t *testing.T) {
	var tests = []itemFieldChangeTest{
		{testname: "A 1 1, decreaseSellIn -1", in: &Item{name: "A", sellIn: 1, quality: 1}, change: -1, want: fmt.Errorf("decreaseSellIn: val is negative, %v", -1)},
		{testname: "A 1 1, decreaseSellIn 1", in: &Item{name: "A", sellIn: 1, quality: 1}, change: 1, want: nil},
	}

	for testInd, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := tt.in.decreaseSellIn(tt.change)
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("test index:%v, got %v, want %v", testInd, err, tt.want)
			}
		})
	}
}
