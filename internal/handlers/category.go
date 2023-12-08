package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amrohan/expenso-go/ent/category"
	"github.com/amrohan/expenso-go/internal/helpers"
	"github.com/go-chi/chi/v5"
)

type Category struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	UserId string `json:"userId"`
}

func (db *Handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	tx, err := db.Category.Query().All(r.Context())

	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Something went wrong", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", tx, nil)
}

func (db *Handler) GetCategoriesByUserId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tx, err := db.Category.Query().Where(category.UserId(id)).All(r.Context())

	if err != nil {
		helpers.SendResponse(w, http.StatusNotFound, "Category Not found", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", tx, nil)
}

func (db *Handler) GetCategoriesByCategoryId(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)

	}
	tx, err := db.Category.Query().Where(category.ID(id)).All(r.Context())

	if err != nil {
		helpers.SendResponse(w, http.StatusNotFound, "Category Not found", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Success", tx, nil)
}

func (db *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid request body", nil, err)
	}
	tx, err := db.Category.Create().SetName(category.Name).SetIcon(category.Icon).SetUserId(category.UserId).Save(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Something went wrong", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Category created successfully", tx, nil)
}

func (db *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
	}
	var category Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid request body", nil, err)
	}
	tx, err := db.Category.UpdateOneID(id).SetName(category.Name).SetIcon(category.Icon).SetUserId(category.UserId).Save(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Something went wrong", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Category updated successfully", tx, nil)
}

func (db *Handler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
	}
	err = db.Category.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusNotFound, "Item does not exist", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Item deleted successfully", nil, nil)
}
