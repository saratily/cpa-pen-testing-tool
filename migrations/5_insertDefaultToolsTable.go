package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table default_tools...")
		_, err := db.Exec(`INSERT INTO TABLE default_tools('type', 'category', 'options', 'format', 'selected')
		VALUES('dig', 'dns_enumeration', '+short', 'dig {url} +short', 't')
		)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("delete from table default_tools...")
		_, err := db.Exec(`DELETE FROM TABLE default_tools`)
		return err
	})
}
