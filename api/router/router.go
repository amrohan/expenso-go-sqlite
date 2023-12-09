package router

import (
	"net/http"

	"github.com/amrohan/expenso-go/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func LoadRoutes(r chi.Router, handler *handlers.Handler) {

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		r.Post("/login", (*handler).Login)
		r.Post("/register", (*handler).Register)
		r.Post("/logout", handlers.Logout)
	})

	r.With(handlers.AuthHandler).Route("/transactions", func(r chi.Router) {
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

	r.With(handlers.AuthHandler).Route("/categories", func(r chi.Router) {
		r.Get("/", handler.GetAllCategories)
		r.Get("/{id}", handler.GetCategoriesByCategoryId)
		r.Get("/user/{id}", handler.GetCategoriesByUserId)
		r.Post("/", handler.CreateCategory)
		r.Put("/{id}", handler.UpdateCategory)
		r.Delete("/{id}", handler.DeleteCategoryById)
	})

	r.With(handlers.AuthHandler).Route("/accounts", func(r chi.Router) {
		r.Get("/", handler.GetAllAccounts)
		r.Get("/{id}", handler.GetAccountByAccountId)
		r.Get("/user/{id}", handler.GetAccountsByUserId)
		r.Post("/", handler.CreateAccount)
		r.Put("/{id}", handler.UpdateAccount)
		r.Delete("/{id}", handler.DeleteAccountById)

	})

	r.With(handlers.AuthHandler).Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUsers)
		r.Get("/{id}", handler.GetUserById)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})

}
