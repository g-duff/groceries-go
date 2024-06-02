package main

import (
	"fmt"

	"github.com/g-duff/groceries/pkg/shoppinglist"
)

func main() {

	meals := [10]string{"chips", "bar"}

	// dd := shoppinglist.MealsDataSpy{}
	dd := shoppinglist.MealsDataFiles{ 
		Directory: "./testdata",
	}

	groceriesListItems := shoppinglist.Create(&dd, meals[:])

	for _, li := range groceriesListItems {
		fmt.Println(li.ToString())
	}

}
