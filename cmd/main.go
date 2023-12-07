package main

import (
	"context"
	"fmt"
	"time"

	"github.com/amrohan/expenso-go/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    client, err := db.ConnectToDB()

    if err != nil {
        panic(err)
    }

    trans,err:= client.Transaction.Create().
    SetTitle("Samosa Pav").
    SetAmount(20).
    SetDate(time.Now()).
    SetType("Expense").
    SetCategoryId(1).
    SetAccountId(1).
    SetCurrency("INR").
    SetUserId("1234").
    Save(context.Background())
    
    if err != nil {
        panic(err)
    }
    
    fmt.Println(trans)
}

