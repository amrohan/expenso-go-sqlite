package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/amrohan/expenso-go/ent"
	"github.com/amrohan/expenso-go/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
			w.Write([]byte("Error decoding"))
			return
		}

		user, err := db.User.Query().Where(user.Email(userCred.Email)).Only(r.Context())
		if err != nil {
			w.Write([]byte("Error querying user"))
			return
		}

		if !CheckPasswordHash(userCred.Password, user.Password) {
			w.Write([]byte("Invalid password"))
			return
		}
		w.Write([]byte("Login successful"))
	}
}

func Register(db *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.Write([]byte("Error decoding"))
			return
		}
		if !user.Validate() {
			w.Write([]byte("Error validating"))
			return
		}
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			w.Write([]byte("Error hashing password"))
			return
		}
		user.Password = hashedPassword
		_, err = db.User.Create().SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).Save(r.Context())
		if err != nil {
			w.Write([]byte("Error saving user"))
			return
		}
		w.Write([]byte("User registered successfully"))
	}
}
