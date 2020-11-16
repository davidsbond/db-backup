// Package backup provides a generic database backup utility. A URL string is used to determine the kind of database
// backup to perform.
package backup

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path/filepath"

	"github.com/davidsbond/db-backup/backup/postgres"
	"github.com/davidsbond/db-backup/backup/sqlite"
)

type (
	// The Backup interface describes types that perform a database backup.
	Backup interface {
		// Do performs a backup, the resulting backup data should be written to the provided
		// io.Writer implementation.
		Do(ctx context.Context, wr io.Writer) error
	}
)

// ErrUnknownScheme is the error returned when a provided url string's scheme does not correspond to an implemented
// backup provider.
var ErrUnknownScheme = errors.New("unknown scheme")

// New returns a new implementation of the Backup interface corresponding to the database provider specified in the
// provided URL string. Returns an error if the database provider cannot be determined.
func New(urlStr string) (Backup, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var user string
	var password string
	if u.User != nil {
		user = u.User.Username()
		password, _ = u.User.Password()
	}

	switch u.Scheme {
	case "postgres":
		return postgres.NewBackup(user, password, u.Host), nil
	case "sqlite":
		return sqlite.NewBackup(filepath.Join("/", u.Host, u.Path)), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownScheme, u.Scheme)
	}
}
