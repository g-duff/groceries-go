package shoppinglist

import (
	"encoding/json"
	"log/slog"
	"os"
	"slices"
	"strings"
)

type MealsData interface {
	getMealsIngredients(meals []string) map[string][]GroceriesListItem
}

type MealsDataSpy struct {
	CountGetMealsIngredients int
}

func (m *MealsDataSpy) getMealsIngredients(meals []string) map[string][]GroceriesListItem {

	m.CountGetMealsIngredients++

	outValue := map[string][]GroceriesListItem{
		"meal1": {
			GroceriesListItem{
				Item:     "rice",
				Quantity: 0.5,
				Unit:     "kg",
			},
			GroceriesListItem{
				Item:     "beans",
				Quantity: 0.2,
				Unit:     "kg",
			},
		},
		"meal2": {
			GroceriesListItem{
				Item:     "rice",
				Quantity: 0.3,
				Unit:     "kg",
			},

			GroceriesListItem{
				Item:     "beans",
				Quantity: 0.1,
				Unit:     "kg",
			},
		},
	}

	return outValue
}

type MealsDataFiles struct { 
	Directory string
}

func (m *MealsDataFiles) getMealsIngredients(meals []string) map[string][]GroceriesListItem {

	data := make(map[string][]GroceriesListItem)

	res, err := os.ReadDir(m.Directory)

	if err != nil {
		slog.Warn("No data returned; directory not found.")
		return data
	}


	// Select all overlapping files in dir entry
	for _, r := range(res) {
		fname := r.Name()
		mealname, found := strings.CutSuffix(fname, ".json")

		if found && slices.Contains(meals, mealname) {
			indata, _ := os.ReadFile(m.Directory + "/"+r.Name())

			var ls []GroceriesListItem
			err = json.Unmarshal(indata, &ls)

			data[mealname] = ls
		}

	}
	return data
}
