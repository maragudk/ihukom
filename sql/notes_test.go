package sql_test

import (
	"context"
	"testing"

	"github.com/maragudk/is"

	"github.com/maragudk/ihukom/model"
	"github.com/maragudk/ihukom/sqltest"
)

func TestDatabase_CRUDNote(t *testing.T) {
	t.Run("can create, read, update, and delete a note", func(t *testing.T) {
		db := sqltest.CreateDatabase(t)
		n, err := db.CreateNote(context.Background())
		is.NotError(t, err)
		is.Equal(t, n.Content, "")

		n.Content = "Yo"

		err = db.SaveNote(context.Background(), n)
		is.NotError(t, err)

		n, err = db.GetNote(context.Background(), n.ID)
		is.NotError(t, err)
		is.Equal(t, n.Content, "Yo")

		err = db.DeleteNote(context.Background(), n.ID)
		is.NotError(t, err)

		n, err = db.GetNote(context.Background(), n.ID)
		is.Error(t, model.ErrorNoteNotFound, err)
	})
}
