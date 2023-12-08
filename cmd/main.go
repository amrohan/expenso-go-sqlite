package main

import (
	"fmt"
	"net/http"

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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/login", handlers.Login(client))
	r.Post("/register", handlers.Register(client))

	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", handler.GetAllTransaction)
		r.Get("/{id}", handler.GetTransactionByTransactionId)
		r.Get("/user/{id}", handler.GetTransactionByUserId)
		r.Get("/category/{id}", handler.GetTransactionsByCategoryId)
		r.Get("/account/{id}", handler.GetTransactionsByAccountId)
		r.Get("/{month}-{year}", handler.GetTransactionBasedOnMonthandYear)
		r.Post("/", handler.CreateTransaction)
		r.Put("/{id}", handler.UpdateTransaction)
		r.Delete("/{id}", handler.DeleteTransaction)
	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", handler.GetAllCategories)
		r.Get("/{id}", handler.GetCategoriesByCategoryId)
		r.Get("/user/{id}", handler.GetCategoriesByUserId)
		r.Post("/", handler.CreateCategory)
		r.Put("/{id}", handler.UpdateCategory)
		r.Delete("/{id}", handler.DeleteCategoryById)
	})

	r.Route("/accounts", func(r chi.Router) {
		r.Get("/", handler.GetAllAccounts)
		r.Get("/{id}", handler.GetAccountByAccountId)
		r.Get("/user/{id}", handler.GetAccountsByUserId)
		r.Post("/", handler.CreateAccount)
		r.Put("/{id}", handler.UpdateAccount)
		r.Delete("/{id}", handler.DeleteAccountById)

	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUsers)
		r.Get("/{id}", handler.GetUserById)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})

	fmt.Println("Server is started")
	http.ListenAndServe(":3000", r)

}
