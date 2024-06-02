package shoppinglist

import (
	"fmt"
	"testing"
)

func TestBusiness(t *testing.T) {

	// Given
	md := MealsDataSpy{CountGetMealsIngredients: 0}
	meals := [10]string{"foo", "bar"}
	expected := []GroceriesListItem{
		{
			Item:     "beans",
			Quantity: 0.1,
			Unit:     "kg",
		},
		{
			Item:     "beans",
			Quantity: 0.2,
			Unit:     "kg",
		},
		{
			Item:     "rice",
			Quantity: 0.3,
			Unit:     "kg",
		},
		{
			Item:     "rice",
			Quantity: 0.5,
			Unit:     "kg",
		},
	}

	// When
	actual := Create(&md, meals[:])

	// Then
	fmt.Println(md)

	if md.CountGetMealsIngredients != 1 {
		t.Fatalf(`getMealsIngredients calls expected: %d, actual: %d`, 1, md.CountGetMealsIngredients)
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Fatalf(`Have: %s, want: %s`, actual[i].ToString(), expected[i].ToString())
		}
	}

}
