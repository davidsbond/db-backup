package main

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/davidsbond/db-backup/backup"
	"pkg.dsb.dev/app"
	"pkg.dsb.dev/closers"
	"pkg.dsb.dev/flag"
	"pkg.dsb.dev/storage/blob"
)

func main() {
	a := app.New(
		app.WithRunner(run),
		app.WithFlags(
			&flag.String{
				Name:        "db-dsn",
				Usage:       "DSN for connecting to the database",
				EnvVar:      "DB_DSN",
				Destination: &dbDSN,
				Required:    true,
			},
			&flag.String{
				Name:        "bucket-dsn",
				Usage:       "DSN for the bucket to place the backup",
				EnvVar:      "BUCKET_DSN",
				Destination: &bucketDSN,
				Required:    true,
			},
			&flag.String{
				Name:        "bucket-dir",
				Usage:       "Location in the bucket to place the backup",
				EnvVar:      "BUCKET_DIR",
				Destination: &bucketDir,
				Required:    true,
			},
		),
	)

	if err := a.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

var (
	dbDSN     string
	bucketDSN string
	bucketDir string
)

func run(ctx context.Context) error {
	bucket, err := blob.OpenBucket(ctx, bucketDSN)
	if err != nil {
		return err
	}
	defer closers.Close(bucket)

	runner, err := backup.New(dbDSN)
	if err != nil {
		return err
	}

	key := filepath.Join(bucketDir, time.Now().Format("2006-01-02.sql.gz"))
	wr, err := bucket.NewWriter(ctx, key)
	if err != nil {
		return err
	}
	defer closers.Close(wr)

	gzipWr := gzip.NewWriter(wr)
	defer closers.Close(gzipWr)

	return runner.Do(ctx, gzipWr)
}
