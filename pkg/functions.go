package pkg

import (
	"fmt"
)

const MAX_LIST_LEN = 100;

func Groceries(meals []string) [MAX_LIST_LEN]GroceriesListItem {
	return makeTestData()
}

type GroceriesListItem struct {
	item string
	quantity float64
	unit string
}


func (g *GroceriesListItem) ToString() string {
	return fmt.Sprintf("%s: %f %s", g.item, g.quantity, g.unit)
}

func makeTestData() [MAX_LIST_LEN]GroceriesListItem {
	// Test data
	myItem := GroceriesListItem{
		item: "rice",
		quantity: 100.0,
		unit: "kg",
	}

	var myList [MAX_LIST_LEN]GroceriesListItem

	myList[0] = myItem
	return myList
}
