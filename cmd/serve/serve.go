package main

import (
	"log/slog"
	"net/http"

	"github.com/g-duff/groceries/pkg/shoppinglist"
)


func main() {

	slog.Info("Starting...")


	http.HandleFunc("/shoppinglist", handle)


	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {

	meals := [10]string{"chips", "bar"}

	dd := shoppinglist.MealsDataSpy{}
	// dd := shoppinglist.MealsDataFiles{ 
	// 	Directory: "./testdata",
	// }

	groceriesListItems := shoppinglist.Create(&dd, meals[:])

	for _, g := range(groceriesListItems) {
		slog.Info(g.ToString())
	}
}
