package db

import (
	"context"
	"log"

	"github.com/amrohan/expenso-go/ent"
)


func ConnectToDB() (*ent.Client,error) {
	client, err := ent.Open("sqlite3", "file:./users.db?_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }

    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
	return client,err
}