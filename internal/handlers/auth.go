package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/amrohan/expenso-go/ent"
	"github.com/amrohan/expenso-go/ent/user"
	"github.com/amrohan/expenso-go/internal/helpers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.Claims
}

var jwtKey = []byte("secret")

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

func (db *Handler) Login(w http.ResponseWriter, r *http.Request) {
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

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: user.Username,
		Claims: jwt.MapClaims{
			"id":  user.ID,
			"exp": expirationTime.Unix(),
			"iat": time.Now().Unix(),
			"sub": user.Username,
			"iss": "expenso-go",
			"aud": "expenso-go",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, "Error signing token", nil, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})
	helpers.SendResponse(w, http.StatusOK, "Login successful", nil, nil)
}

func (db *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	id := uuid.New()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid body", nil, err)
		return
	}
	if !user.Validate() {
		helpers.SendResponse(w, http.StatusBadRequest, "Please send valid body", nil, err)
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
		// use switch case to check if username already exists or not
		switch err.(type) {
		case *ent.ConstraintError:
			if strings.Contains(err.Error(), "users.email") {
				helpers.SendResponse(w, http.StatusBadRequest, "Email already exists", nil, err)
				return
			}
			helpers.SendResponse(w, http.StatusBadRequest, "Username already exists", nil, err)

			return
		default:
			helpers.SendResponse(w, http.StatusInternalServerError, "Error creating user", nil, err)
			return
		}
	}

	helpers.SendResponse(w, http.StatusOK, "User created successfully", nil, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})
	helpers.SendResponse(w, http.StatusOK, "Logout successful", nil, nil)
}

// func RefreshToken(db *ent.Client) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				helpers.SendResponse(w, http.StatusUnauthorized, "Unauthorized", nil, err)
// 				return
// 			}
// 			helpers.SendResponse(w, http.StatusBadRequest, "Bad request", nil, err)
// 			return
// 		}
// 		tokenStr := cookie.Value

// 		claims := &Claims{}

// 		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})
// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				helpers.SendResponse(w, http.StatusUnauthorized, "Unauthorized", nil, err)
// 				return
// 			}
// 			helpers.SendResponse(w, http.StatusBadRequest, "Bad request", nil, err)
// 			return
// 		}
// 		if !tkn.Valid {
// 			helpers.SendResponse(w, http.StatusUnauthorized, "Unauthorized", nil, err)
// 			return
// 		}

// 		// Token is valid, generate a new one with a new expiration time
// 		expirationTime := time.Now().Add(time.Minute * 5)

// 	}
// }

type userKey string

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the token from the cookie
		token, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				helpers.SendResponse(w, http.StatusUnauthorized, "Unauthorized", nil, err)
				return
			}
			helpers.SendResponse(w, http.StatusBadRequest, "Bad request", nil, err)
			return
		}
		tokenStr := token.Value
		// parse the token
		claims := &Claims{}
		tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				helpers.SendResponse(w, http.StatusUnauthorized, "Unauthorized", nil, err)
				return
			}
			helpers.SendResponse(w, http.StatusBadRequest, "Bad request", nil, err)
			return
		}
		if !tkn.Valid {
			helpers.SendResponse(w, http.StatusUnauthorized, "Unauthorized", nil, err)
			return
		}
		// set the user context
		ctx := context.WithValue(r.Context(), userKey("user"), claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
