package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table penetrations...")
		_, err := db.Exec(`CREATE TABLE penetrations(
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			website TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMPTZ DEFAULT NULL,
			user_id INT REFERENCES users ON DELETE CASCADE
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table penetrations...")
		_, err := db.Exec(`DROP TABLE penetrations`)
		return err
	})
}
