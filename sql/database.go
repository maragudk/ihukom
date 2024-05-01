package sql

import (
	"context"
	"embed"
	"io/fs"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/maragudk/migrate"
	"github.com/maragudk/snorkel"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB   *sqlx.DB
	log  *snorkel.Logger
	path string
}

type NewDatabaseOptions struct {
	Log  *snorkel.Logger
	Path string
}

// NewDatabase with the given options.
// If no logger is provided, logs are discarded.
func NewDatabase(opts NewDatabaseOptions) *Database {
	// - Set WAL mode (not strictly necessary each time because it's persisted in the database, but good for first run)
	// - Set busy timeout, so concurrent writers wait on each other instead of erroring immediately
	// - Enable foreign key checks
	opts.Path += "?_journal=WAL&_timeout=5000&_fk=true"

	return &Database{
		path: opts.Path,
		log:  opts.Log,
	}
}

func (d *Database) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	d.log.Event("Starting database", 1, "path", d.path)

	var err error
	d.DB, err = sqlx.ConnectContext(ctx, "sqlite3", d.path)
	if err != nil {
		return err
	}

	return nil
}

//go:embed migrations
var migrations embed.FS

func (d *Database) MigrateUp(ctx context.Context) error {
	fsys := d.getMigrations()
	return migrate.Up(ctx, d.DB.DB, fsys)
}

func (d *Database) MigrateDown(ctx context.Context) error {
	fsys := d.getMigrations()
	return migrate.Down(ctx, d.DB.DB, fsys)
}

func (d *Database) getMigrations() fs.FS {
	fsys, err := fs.Sub(migrations, "migrations")
	if err != nil {
		panic(err)
	}
	return fsys
}

func (d *Database) Ping(ctx context.Context) error {
	return d.DB.PingContext(ctx)
}
