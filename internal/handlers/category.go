package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amrohan/expenso-go/ent"
	"github.com/amrohan/expenso-go/ent/category"
	"github.com/go-chi/chi/v5"
)


type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	UserId      string `json:"userId"`
}


func GetAllCategories(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		tx, err := db.Category.Query().All(r.Context())
		
		if err !=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func GetCategoriesByUserId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		id := chi.URLParam(r, "id")
		tx, err := db.Category.Query().Where(category.UserId(id)).All(r.Context())
		
		if err !=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func GetCategoriesByCategoryId(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		strId := chi.URLParam(r, "id")
		id,err :=strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Fail to convert Id"))
		}
		tx, err := db.Category.Query().Where(category.ID(id)).All(r.Context())
		
		if err !=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func CreateCategory(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var category Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Fail to decode request body"))
		}
		tx, err := db.Category.Create().SetName(category.Name).SetIcon(category.Icon).SetUserId(category.UserId).Save(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func UpdateCategory(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		strId := chi.URLParam(r, "id")
		id,err :=strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Fail to convert Id"))
		}
		var category Category
		err = json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Fail to decode request body"))
		}
		tx, err := db.Category.UpdateOneID(id).SetName(category.Name).SetIcon(category.Icon).SetUserId(category.UserId).Save(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	}
}

func DeleteCategoryById(db *ent.Client) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		strId := chi.URLParam(r, "id")
		id,err :=strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Fail to convert Id"))
		}
		err = db.Category.DeleteOneID(id).Exec(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}