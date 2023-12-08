package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amrohan/expenso-go/ent/account"
	"github.com/amrohan/expenso-go/internal/helpers"
	"github.com/amrohan/expenso-go/internal/models"
	"github.com/go-chi/chi/v5"
)

func (db *Handler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := db.Account.Query().All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", accounts, nil)
}

func (db *Handler) GetAccountsByUserId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	accounts, err := db.Account.Query().Where(account.UserIdEQ(id)).All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", accounts, nil)
}

func (db *Handler) GetAccountByAccountId(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}
	account, err := db.Account.Query().Where(account.ID(id)).All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", account, nil)
}

func (db *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newAccount, err := db.Account.Create().SetName(account.Name).SetIcon(account.Icon).SetUserId(account.UserId).Save(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", newAccount, nil)

}

func (db *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid body", nil, err)
		return
	}
	updatedAccount, err := db.Account.UpdateOneID(id).SetName(account.Name).SetIcon(account.Icon).SetUserId(account.UserId).Save(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error", nil, err)

		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", updatedAccount, nil)

}

func (db *Handler) DeleteAccountById(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}
	err = db.Account.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusNotFound, "Error", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Account deleted successfully", nil, nil)

}
