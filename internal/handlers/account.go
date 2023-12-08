package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amrohan/expenso-go/ent"
	"github.com/amrohan/expenso-go/ent/account"
	"github.com/go-chi/chi/v5"
)


type Account struct{
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	UserId      string `json:"userId"`
}


func GetAllAccounts(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		accounts, err := db.Account.Query().All(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(accounts); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func GetAccountsByUserId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		id := r.URL.Query().Get("id")
		accounts, err := db.Account.Query().Where(account.UserIdEQ(id)).All(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(accounts); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetAccountByAccountId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		strId := chi.URLParam(r, "id")
	    id, err := strconv.Atoi(strId); 
	    if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		account, err := db.Account.Query().Where(account.ID(id)).All(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(account); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CreateAccount(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var account Account
		if err := json.NewDecoder(r.Body).Decode(&account); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newAccount, err := db.Account.Create().SetName(account.Name).SetIcon(account.Icon).SetUserId(account.UserId).Save(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(newAccount); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateAccount(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		strId := chi.URLParam(r, "id")
	    id, err := strconv.Atoi(strId); 
	    if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var account Account
		if err := json.NewDecoder(r.Body).Decode(&account); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedAccount, err := db.Account.UpdateOneID(id).SetName(account.Name).SetIcon(account.Icon).SetUserId(account.UserId).Save(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(updatedAccount); err != nil{
		
		http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteAccountById(db *ent.Client)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		strId := chi.URLParam(r, "id")
	    id, err := strconv.Atoi(strId); 
	    if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.Account.DeleteOneID(id).Exec(r.Context())
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Account deleted successfully"})
	
	}
}