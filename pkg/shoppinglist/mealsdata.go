package shoppinglist

type MealsData interface {
	getMealsIngredients(meals []string) map[string][]GroceriesListItem
}

type SpyMealsData struct {
	CountGetMealsIngredients int
}

func (m *SpyMealsData) getMealsIngredients(meals []string) map[string][]GroceriesListItem {

	m.CountGetMealsIngredients++

	outValue := map[string][]GroceriesListItem{
		"meal1": {
			GroceriesListItem{
				item:     "rice",
				quantity: 0.5,
				unit:     "kg",
			},
			GroceriesListItem{
				item:     "beans",
				quantity: 0.2,
				unit:     "kg",
			},
		},
		"meal2": {
			GroceriesListItem{
				item:     "rice",
				quantity: 0.3,
				unit:     "kg",
			},

			GroceriesListItem{
				item:     "beans",
				quantity: 0.1,
				unit:     "kg",
			},
		},
	}

	return outValue
}
