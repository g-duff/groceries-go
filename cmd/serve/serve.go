package main

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/g-duff/groceries/pkg/shoppinglist"
)

func main() {
	slog.Info("Starting...")

	srvExit := make(chan os.Signal, 1)
	srvErr := make(chan error, 1)

	signal.Notify(srvExit, os.Interrupt)

	srv := http.Server{
		Addr:    ":8080",
		Handler: newHandler(),
	}

	go func() {
		srvErr <- srv.ListenAndServe()
		slog.Info("Server stopped listening for new connections")
	}()

	select {
	case err := <-srvErr:
		slog.Error(err.Error())
		return
	case <-srvExit:
		slog.Info("Server shutting down...")
		srv.Shutdown(context.Background())
	}
	slog.Info("Server has shut down")
}

func newHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/shoppinglist", handleShoppingList)

	return mux
}

func handleShoppingList(w http.ResponseWriter, r *http.Request) {

	meals := [10]string{"chips", "bar"}

	// dd := shoppinglist.MealsDataSpy{}
	dd := shoppinglist.MealsDataFiles{
		Directory: "./testdata",
	}

	groceriesListItems := shoppinglist.Create(&dd, meals[:])

	resp, err := json.Marshal(groceriesListItems)

	if err != nil {
		io.WriteString(w, "{ \"message\": \"request failed\" }")
	}

	io.WriteString(w, string(resp))
}
