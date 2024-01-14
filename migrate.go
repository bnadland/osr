package main

import (
	"net/url"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
)

func Migrate(databaseUrl string) error {
	u, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	d := dbmate.New(u)
	d.FS = migrations
	d.AutoDumpSchema = false
	return d.CreateAndMigrate()
}
