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
		if myList[i].item != myList[j].item {
			return myList[i].item < myList[j].item
		}
		return myList[i].quantity < myList[j].quantity

	})

	return toSort
}

type GroceriesListItem struct {
	item     string
	quantity float64
	unit     string
}

func (g *GroceriesListItem) ToString() string {
	return fmt.Sprintf("%s: %f %s", g.item, g.quantity, g.unit)
}
