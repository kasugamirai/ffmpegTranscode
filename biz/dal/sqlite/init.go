package sqlite

import (
	"context"
	"entgo.io/ent/dialect"
	"freefrom.space/videoTransform/conf"
	"freefrom.space/videoTransform/ent"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Client is a global variable to access the ent.Client
var Client *ent.Client

func Init() {
	var err error
	// Get the configuration
	c := conf.GetConf()

	Client, err = ent.Open(dialect.SQLite, c.Sqlite.DSN)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
