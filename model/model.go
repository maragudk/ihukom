package model

import (
	"database/sql/driver"
	"time"

	"github.com/maragudk/errors"
)

type Time struct {
	T time.Time
}

// RFC3339Milli is like time.RFC3339Nano, but with millisecond precision, and fractional seconds do not have trailing
// zeros removed.
const RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"

// Value satisfies driver.Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return t.T.UTC().Format(RFC3339Milli), nil
}

// Scan satisfies sql.Scanner interface.
func (t *Time) Scan(src any) error {
	if src == nil {
		return nil
	}

	s, ok := src.(string)
	if !ok {
		return errors.Newf("error scanning time, got %+v", src)
	}

	parsedT, err := time.Parse(RFC3339Milli, s)
	if err != nil {
		return err
	}

	t.T = parsedT.UTC()

	return nil
}

type ID string

func (i ID) String() string {
	return string(i)
}

type Note struct {
	ID      ID
	Created Time
	Updated Time
	Content string
}
