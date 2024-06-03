package main

import (
	"encoding/json"
	"io"
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

	resp, err := json.Marshal(groceriesListItems)

	if err != nil {
		io.WriteString(w, "{ \"message\": \"request failed\" }")
	}

	io.WriteString(w, string(resp))
}
