package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lukas-blaha/quickle/pkg/formats"
)

func GetLesson(db *sql.DB, lesson string) ([]formats.Entry, error) {
	var data []formats.Entry
	var e formats.Entry
	rows, err := db.Query(fmt.Sprintf(`SELECT "term", "def" FROM %s`, lesson))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Term, &e.Def)
		if err != nil {
			return nil, err
		}

		data = append(data, e)
	}

	return data, nil
}

func GetStudySets(db *sql.DB) ([]string, error) {
	var name string
	var lessons []string
	rows, err := db.Query(`
			SELECT
				table_name
			FROM information_schema.tables
			WHERE
				table_schema != 'information_schema' and table_schema != 'pg_catalog';
	`)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, name)
	}

	return lessons, nil
}
