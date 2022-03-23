package main

import (
	"reflect"
	"testing"
)

func TestUpdateQuality(t *testing.T) {
	var tests = []struct {
		testname string
		in       []*Item
		want     []*Item
	}{
		{testname: "A 1 1", in: []*Item{{name: "A", sellIn: 1, quality: 1}}, want: []*Item{{name: "A", sellIn: 0, quality: 0}}},
		{testname: "A 1 0", in: []*Item{{name: "A", sellIn: 1, quality: 0}}, want: []*Item{{name: "A", sellIn: 0, quality: 0}}},
		{testname: "Sulfuras, Hand of Ragnaros 0 1", in: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 1}}, want: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 1}}},
		{testname: "Aged Brie 1 50", in: []*Item{{name: "Aged Brie", sellIn: 1, quality: 50}}, want: []*Item{{name: "Aged Brie", sellIn: 0, quality: 50}}},
		{testname: "Aged Brie 1 49", in: []*Item{{name: "Aged Brie", sellIn: 1, quality: 49}}, want: []*Item{{name: "Aged Brie", sellIn: 0, quality: 50}}},
		{testname: "Backstage passes to a TAFKAL80ETC concert 20 49", in: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 20, quality: 49}}, want: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 19, quality: 50}}},
		{testname: "Backstage passes to a TAFKAL80ETC concert 5 49", in: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 49}}, want: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50}}},
		{testname: "Backstage passes to a TAFKAL80ETC concert 5 40", in: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 40}}, want: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 43}}},
		{testname: "Backstage passes to a TAFKAL80ETC concert 10 40", in: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 40}}, want: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 42}}}, //end of first section
		{testname: "Sulfuras, Hand of Ragnaros -1 1", in: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: -1, quality: 1}}, want: []*Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: -1, quality: 1}}},
		{testname: "Aged Brie 0 40", in: []*Item{{name: "Aged Brie", sellIn: 0, quality: 40}}, want: []*Item{{name: "Aged Brie", sellIn: -1, quality: 42}}},
		{testname: "Backstage passes to a TAFKAL80ETC concert 0 10", in: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 0, quality: 10}}, want: []*Item{{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: -1, quality: 0}}},
		{testname: "A 0 10", in: []*Item{{name: "A", sellIn: 0, quality: 10}}, want: []*Item{{name: "A", sellIn: -1, quality: 8}}},
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
