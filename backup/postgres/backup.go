// Package postgres contains an implementation of backup.Backup that supports postgres databases.
package postgres

import (
	"context"
	"io"
	"os"
	"os/exec"
)

type (
	// The Backup type is an implementation of backup.Backup that creates dumps of postgres databases.
	Backup struct {
		user     string
		password string
		host     string
	}
)

// NewBackup returns a new instance of the Backup type that will perform a backup against the provided
// postgres database.
func NewBackup(user, password, host string) *Backup {
	return &Backup{
		user:     user,
		password: password,
		host:     host,
	}
}

// Do performs a backup on the configured postgres database and writes the contents to the provided
// io.Writer implementation.
func (b *Backup) Do(_ context.Context, wr io.Writer) error {
	cmd := exec.Command("pg_dumpall")
	cmd.Env = append(os.Environ(),
		"PGUSER="+b.user,
		"PGPASSWORD="+b.password,
		"PGHOST="+b.host,
	)

	cmd.Stdout = wr
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
