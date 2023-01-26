package web

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"hello-go-todo-app/db"
	"net/http"
)

func Migrate(w http.ResponseWriter, r *http.Request) {
	driver, err := postgres.WithInstance(db.GetDB(), &postgres.Config{})

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migration", "postgres", driver)

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}

	m.Up()

	fmt.Fprintf(w, "Migrated")
}
