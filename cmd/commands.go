package main

import (
	"fmt"

	"github.com/g-duff/groceries/pkg"
)

func main() {

	meals := [10]string{"foo", "bar"}

	groceriesListItems := pkg.Groceries(meals[:])

	for _, li := range groceriesListItems {
		if li != (pkg.GroceriesListItem{}) {
			fmt.Println(li.ToString())
		}
	}

}
