package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/amrohan/expenso-go/ent"
	"github.com/amrohan/expenso-go/ent/transaction"
	"github.com/go-chi/chi/v5"
)

type Transaction struct{
	Id int `json:"id"`
	UserId string `json:"user_id"`
	Title string `json:"title"`
	Amount float64 `json:"amount"`
	Description string `json:"description"`
	Date time.Time `json:"date"`
	Type string `json:"type"`
	CategoryId int `json:"category_id"`
	AccountId int `json:"account_id"`
	Status string `json:"status"`
	Currency string `json:"currency"`
}



func CreateTransaction(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var transaction Transaction

		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tx, err:= db.Transaction.
			Create().
			SetTitle(transaction.Title).
			SetAmount(transaction.Amount).
			SetDescription(transaction.Description).
			SetDate(transaction.Date).
			SetCategoryId(transaction.CategoryId).
			SetAccountId(transaction.AccountId).
			SetUserId(transaction.UserId).
			SetType(transaction.Type).
			SetStatus(transaction.Status).
			SetCurrency(transaction.Currency).
			Save(r.Context())

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(tx)
		
	}
}

func GetAllTransaction(db *ent.Client)http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request){
		tx, err := db.Transaction.Query().All(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func GetTransactionByTransactionId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r,"id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tx, err := db.Transaction.Query().Where(transaction.ID(id)).Only(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func GetTransactionByUserId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r,"id")
		tx, err := db.Transaction.Query().Where(transaction.UserId(id)).All(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func GetTransactionsByCategoryId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r,"id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tx, err := db.Transaction.Query().Where(transaction.CategoryId(id)).All(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func GetTransactionsByAccountId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r,"id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tx, err := db.Transaction.Query().Where(transaction.AccountId(id)).All(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func GetTransactionBasedOnMonthandYear(db *ent.Client)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		monthStr := chi.URLParam(r,"month")
		yearStr := chi.URLParam(r,"year")

		month, err := strconv.Atoi(monthStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		endDate := startDate.AddDate(0, 1, 0)

		tx, err := db.Transaction.Query().Where(transaction.DateGTE(startDate), transaction.DateLT(endDate)).All(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func UpdateTransaction(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idStr := chi.URLParam(r,"id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var transaction Transaction

		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tx, err := db.Transaction.UpdateOneID(id).
			SetTitle(transaction.Title).
			SetAmount(transaction.Amount).
			SetDescription(transaction.Description).
			SetDate(transaction.Date).
			SetCategoryId(transaction.CategoryId).
			SetAccountId(transaction.AccountId).
			SetUserId(transaction.UserId).
			SetType(transaction.Type).
			SetStatus(transaction.Status).
			SetCurrency(transaction.Currency).
			Save(r.Context())

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(tx)
		
	}
}

func DeleteTransaction(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idStr := chi.URLParam(r,"id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = db.Transaction.DeleteOneID(id).Exec(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Item deleted")
	}
}