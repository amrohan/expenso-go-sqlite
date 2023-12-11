package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/amrohan/expenso-go/api/router"
	"github.com/amrohan/expenso-go/internal/db"
	"github.com/amrohan/expenso-go/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	client, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	handler := handlers.NewHandler(client)

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.CleanPath)

	router.LoadRoutes(r, handler)

	fmt.Println("Server is started")

	// server is running on
	fmt.Println("Server is running on port " + port)
	http.ListenAndServe(":"+port, r)

}
