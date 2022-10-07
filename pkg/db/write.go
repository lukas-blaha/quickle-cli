package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/lukas-blaha/quickle/pkg/formats"
)

func NewTable(db *sql.DB, name string) error {
	_, err := db.Query(fmt.Sprintf(`CREATE TABLE %s (
											term varchar(3000),
											def varchar(3000)
	);`, name))
	if err != nil {
		return err
	}

	return nil
}

func ImportData(db *sql.DB, name string, entries []formats.Entry) {
	for _, v := range entries {
		_, err := db.Query(fmt.Sprintf(
			`INSERT INTO %s (term, def) VALUES ('%s', '%s');`,
			name, v.Term, v.Def))
		if err != nil {
			log.Fatal(err)
		}
	}
}
