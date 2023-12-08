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

    r := chi.NewRouter()
    r.Use(middleware.Logger,middleware.CleanPath)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

    r.Post("/login", handlers.Login(client))
    r.Post("/register", handlers.Register(client))
    
    r.Group(func(r chi.Router) {
        r.Get("/transactions", handlers.GetAllTransaction(client))
        r.Get("/transactions/{id}", handlers.GetTransactionByTransactionId(client))
        r.Get("/transactions/user/{id}", handlers.GetTransactionByUserId(client))
        r.Get("/transactions/category/{id}", handlers.GetTransactionsByCategoryId(client))
        r.Get("/transactions/account/{id}", handlers.GetTransactionsByAccountId(client))
        r.Get("/transactions/{month}-{year}", handlers.GetTransactionBasedOnMonthandYear(client))
        r.Post("/transactions", handlers.CreateTransaction(client))
        r.Put("/transactions/{id}", handlers.UpdateTransaction(client))
        r.Delete("/transactions/{id}", handlers.DeleteTransaction(client))
    })

    r.Group(func(r chi.Router) {
        r.Get("/categories",handlers.GetAllCategories(client))
        r.Get("/categories/{id}",handlers.GetCategoriesByCategoryId(client))
        r.Get("/categories/user/{id}",handlers.GetCategoriesByUserId(client))
        r.Post("/categories",handlers.CreateCategory(client))
        r.Put("/categories/{id}",handlers.UpdateCategory(client))
        r.Delete("/categories/{id}",handlers.DeleteCategoryById(client))
    })

    r.Group(func(r chi.Router) {
        r.Get("/accounts",handlers.GetAllAccounts(client))
        r.Get("/accounts/{id}",handlers.GetAccountByAccountId(client))
        r.Get("/accounts/user/{id}",handlers.GetAccountsByUserId(client))
        r.Post("/accounts",handlers.CreateAccount(client))
        r.Put("/accounts/{id}",handlers.UpdateAccount(client))
        r.Delete("/accounts/{id}",handlers.DeleteAccountById(client))
    })

    fmt.Println("Server is started")
    http.ListenAndServe(":3000", r)

}

