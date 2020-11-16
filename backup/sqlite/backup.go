// Package sqlite contains an implementation of backup.Backup that supports sqlite databases.
package sqlite

import (
	"context"
	"io"
	"os"
	"os/exec"
)

type (
	// The Backup type is an implementation of backup.Backup that creates dumps of sqlite databases.
	Backup struct {
		file string
	}
)

// NewBackup returns a new instance of the Backup type that will perform a backup against the provided
// file. It is expected that the file is an sqlite database.
func NewBackup(file string) *Backup {
	return &Backup{
		file: file,
	}
}

// Do performs a backup on the configured sqlite database and writes the contents to the provided
// io.Writer implementation.
func (b *Backup) Do(_ context.Context, wr io.Writer) error {
	cmd := exec.Command("sqlite3", b.file, ".dump")

	cmd.Stdout = wr
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
