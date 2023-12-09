package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/amrohan/expenso-go/ent"
	"github.com/amrohan/expenso-go/ent/user"
	"github.com/amrohan/expenso-go/internal/helpers"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *User) Validate() bool {
	return u.Username != "" && u.Email != "" && u.Password != ""
}

func Login(db *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userCred User

		err := json.NewDecoder(r.Body).Decode(&userCred)
		if err != nil {
			helpers.SendResponse(w, http.StatusBadRequest, "Please send valid body", nil, err)
			return
		}

		user, err := db.User.Query().Where(user.Email(userCred.Email)).Only(r.Context())
		if err != nil {
			helpers.SendResponse(w, http.StatusBadRequest, "Invalid email", nil, err)
			return
		}

		if !CheckPasswordHash(userCred.Password, user.Password) {
			helpers.SendResponse(w, http.StatusBadRequest, "Invalid password", nil, err)
			return
		}
		helpers.SendResponse(w, http.StatusOK, "Login successful", nil, nil)
	}
}

func Register(db *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		id := uuid.New()
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			helpers.SendResponse(w, http.StatusBadRequest, "Please send valid body", nil, err)
			return
		}
		if !user.Validate() {
			w.Write([]byte("Error validating"))
			return
		}
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			helpers.SendResponse(w, http.StatusInternalServerError, "Error hashing password", nil, err)
			return
		}
		user.Password = hashedPassword
		_, err = db.User.Create().SetID(id).SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).Save(r.Context())
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		helpers.SendResponse(w, http.StatusOK, "User created successfully", nil, nil)
	}
}
