package models

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Account struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	UserId string `json:"userId"`
}

type Transaction struct {
	Id          int       `json:"id"`
	UserId      string    `json:"user_id"`
	Title       string    `json:"title"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"`
	CategoryId  int       `json:"category_id"`
	AccountId   int       `json:"account_id"`
	Status      string    `json:"status"`
	Currency    string    `json:"currency"`
}
