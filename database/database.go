package database

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/fs/ent"
	"entgo.io/ent/examples/fs/ent/migrate"
	"github.com/HDYS-TTBYS/go-todo-api/config"
	_ "github.com/lib/pq"
)

var client *ent.Client

// Init initializes database
func Init(isReset bool, models ...interface{}) {
	c := config.GetConfig()
	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		c.PGHOST, c.PGUSER, c.PGPASSWORD, c.PGDATABASE, c.PGPORT,
	),
	)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// マイグレーションの実行
	if isReset {
		err = client.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}
}

func GetClient() *ent.Client {
	return client
}
