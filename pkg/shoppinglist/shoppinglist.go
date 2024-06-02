package shoppinglist

import (
	"fmt"
	"sort"
)

const MAX_LIST_LEN = 100

func Create(md MealsData, meals []string) []GroceriesListItem {

	mm := md.getMealsIngredients(meals)

	var myList [MAX_LIST_LEN]GroceriesListItem

	outer_count := 0
	for _, meal := range mm {
		for _, thg := range meal {
			if thg != (GroceriesListItem{}) {
				myList[outer_count] = thg
				outer_count++
			}
		}
	}

	toSort := myList[:outer_count]

	sort.Slice(toSort, func(i int, j int) bool {
		if myList[i].Item != myList[j].Item {
			return myList[i].Item < myList[j].Item
		}
		return myList[i].Quantity < myList[j].Quantity

	})

	return toSort
}

type GroceriesListItem struct {
	Item     string
	Quantity float64
	Unit     string
}

func (g *GroceriesListItem) ToString() string {
	return fmt.Sprintf("%s: %f %s", g.Item, g.Quantity, g.Unit)
}
