package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/amrohan/expenso-go/ent/transaction"
	"github.com/amrohan/expenso-go/internal/helpers"
	"github.com/amrohan/expenso-go/internal/models"
	"github.com/go-chi/chi/v5"
)

func (db *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid JSON", nil, err)
		return
	}
	tx, err := db.Transaction.
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
		helpers.SendResponse(w, http.StatusInternalServerError, "Error creating transaction", nil, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Transaction created successfully", tx, nil)

}

func (db *Handler) GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	tx, err := db.Transaction.Query().All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error getting transactions", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Transactions fetched successfully", tx, nil)
}

func (db *Handler) GetTransactionByTransactionId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}

	tx, err := db.Transaction.Query().Where(transaction.ID(id)).Only(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error getting transaction", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Transaction fetched successfully", tx, nil)

}

func (db *Handler) GetTransactionByUserId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tx, err := db.Transaction.Query().Where(transaction.UserId(id)).All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error getting transaction", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Transaction fetched successfully", tx, nil)
}

func (db *Handler) GetTransactionsByCategoryId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}
	tx, err := db.Transaction.Query().Where(transaction.CategoryId(id)).All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error getting transaction", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Transaction fetched successfully", tx, nil)
}

func (db *Handler) GetTransactionsByAccountId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}
	tx, err := db.Transaction.Query().Where(transaction.AccountId(id)).All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error getting transaction", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Transaction fetched successfully", tx, nil)
}

func (db *Handler) GetTransactionBasedOnMonthandYear(w http.ResponseWriter, r *http.Request) {
	monthStr := chi.URLParam(r, "month")
	yearStr := chi.URLParam(r, "year")

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid month", nil, err)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid year", nil, err)
		return
	}

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	tx, err := db.Transaction.Query().Where(transaction.DateGTE(startDate), transaction.DateLT(endDate)).All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error getting transaction", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Transaction fetched successfully", tx, nil)
}

func (db *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}

	var transaction models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid JSON", nil, err)
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
		helpers.SendResponse(w, http.StatusInternalServerError, "Error updating transaction", nil, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "Transaction updated successfully", tx, nil)

}

func (db *Handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}

	err = db.Transaction.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error deleting transaction", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Transaction deleted successfully", nil, nil)
}
