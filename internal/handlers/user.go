package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amrohan/expenso-go/ent"
	"github.com/amrohan/expenso-go/ent/user"
	"github.com/amrohan/expenso-go/internal/helpers"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*ent.Client
}

func NewHandler(client *ent.Client) *Handler {
	return &Handler{Client: client}
}

func (db *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user, err := db.User.Query().All(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "failed", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "success", user, nil)
}

func (db *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid ID", nil, err)
		return
	}
	user, err := db.User.Query().Where(user.ID(id)).Only(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "failed", nil, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, "success", user, nil)
}

func (db *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid data", nil, err)
		return
	}
	u, err := db.User.Create().SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).Save(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "success", u, nil)
}

func (db *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid data", nil, err)
		return
	}

	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid ID", nil, err)
		return
	}

	u, err := db.User.UpdateOneID(id).SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).Save(r.Context())

	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "failed", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "success", u, nil)
}

func (db *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid ID", nil, err)
		return
	}

	err = db.User.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "failed", nil, err)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "Deleted Successfully", nil, nil)
}
