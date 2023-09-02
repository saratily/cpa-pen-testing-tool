package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table pen_tools...")
		_, err := db.Exec(`CREATE TABLE pen_tools(
			id SERIAL PRIMARY KEY,
			unique_id VARCHAR(36) NOT NULL,
			type TEXT NOT NULL,
			category TEXT NOT NULL,
			options TEXT,
			command TEXT,
			output TEXT,
			selected BOOLEAN NOT NULL DEFAULT 't',
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMPTZ DEFAULT NULL,
			penetration_id INT REFERENCES penetrations ON DELETE CASCADE
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table pen_tools...")
		_, err := db.Exec(`DROP TABLE pen_tools`)
		return err
	})
}
