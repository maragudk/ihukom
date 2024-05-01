package sqltest

import (
	"context"
	"io"
	"testing"

	"github.com/maragudk/snorkel"

	"github.com/maragudk/ihukom/sql"
)

// CreateDatabase for testing.
func CreateDatabase(t *testing.T) *sql.Database {
	t.Helper()

	db := sql.NewDatabase(sql.NewDatabaseOptions{
		Log:  snorkel.New(snorkel.Options{W: io.Discard}),
		Path: ":memory:",
	})
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}

	if err := db.MigrateUp(context.Background()); err != nil {
		t.Fatal(err)
	}

	return db
}
