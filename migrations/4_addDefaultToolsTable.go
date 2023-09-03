package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table default_tools...")
		_, err := db.Exec(`CREATE TABLE default_tools(
			id SERIAL PRIMARY KEY,
			type TEXT NOT NULL,
			category TEXT NOT NULL,
			options TEXT,
			format TEXT,
			active BOOLEAN NOT NULL DEFAULT 't',
			can_change BOOLEAN NOT NULL DEFAULT 't',
			selected BOOLEAN NOT NULL DEFAULT 't',
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMPTZ DEFAULT NULL
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table default_tools...")
		_, err := db.Exec(`DROP TABLE default_tools`)
		return err
	})
}
