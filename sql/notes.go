package sql

import (
	"context"
	"database/sql"

	"github.com/maragudk/errors"

	"github.com/maragudk/ihukom/model"
)

func (d *Database) GetNotes(ctx context.Context) ([]model.Note, error) {
	var notes []model.Note
	err := d.DB.SelectContext(ctx, &notes, `select * from notes order by created desc`)
	return notes, err
}

func (d *Database) GetNote(ctx context.Context, id model.ID) (model.Note, error) {
	var n model.Note
	if err := d.DB.GetContext(ctx, &n, `select * from notes where id = ?`, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return n, model.ErrorNoteNotFound
		}
		return n, err
	}
	return n, nil
}

func (d *Database) CreateNote(ctx context.Context) (model.Note, error) {
	var n model.Note
	err := d.DB.GetContext(ctx, &n, `insert into notes (content) values ('') returning *`)
	return n, err
}

func (d *Database) SaveNote(ctx context.Context, n model.Note) error {
	query := `
		insert into notes (id, content) values (?, ?)
		on conflict (id) do update set
			content = excluded.content`
	if _, err := d.DB.ExecContext(ctx, query, n.ID, n.Content); err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteNote(ctx context.Context, id model.ID) error {
	_, err := d.DB.ExecContext(ctx, `delete from notes where id = ?`, id)
	return err
}
