package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table tools...")
		_, err := db.Exec(`CREATE TABLE tools(
			id SERIAL PRIMARY KEY,
			unique_id VARCHAR(36) NOT NULL,
			type TEXT NOT NULL,
			category TEXT NOT NULL,
			options TEXT,
			command TEXT,
			output TEXT,
			can_change BOOLEAN NOT NULL DEFAULT 't',
			selected BOOLEAN NOT NULL DEFAULT 't',
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			penetration_id INT REFERENCES penetrations ON DELETE CASCADE
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table tools...")
		_, err := db.Exec(`DROP TABLE tools`)
		return err
	})
}
