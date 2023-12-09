package main

import (
	"fmt"
	"net/http"

	"github.com/amrohan/expenso-go/api/router"
	"github.com/amrohan/expenso-go/internal/db"
	"github.com/amrohan/expenso-go/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	handler := handlers.NewHandler(client)

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.CleanPath)

	router.LoadRoutes(r, handler)

	fmt.Println("Server is started")
	http.ListenAndServe(":3000", r)

}
