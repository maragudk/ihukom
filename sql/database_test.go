package sql_test

import (
	"context"
	"testing"

	"github.com/maragudk/is"

	"github.com/maragudk/ihukom/sqltest"
)

func TestDatabase_Migrate(t *testing.T) {
	t.Run("can migrate down and back up", func(t *testing.T) {
		db := sqltest.CreateDatabase(t)

		err := db.MigrateDown(context.Background())
		is.NotError(t, err)

		err = db.MigrateUp(context.Background())
		is.NotError(t, err)
	})
}
