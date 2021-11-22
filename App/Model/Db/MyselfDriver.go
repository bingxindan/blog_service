package Db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func NewOpenDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bxd_blog")
	if err != nil {
		panic(err)
	}
	return db
}

func MyselfQueryNoRows(ctx context.Context, id int) error {
	db := NewOpenDb()

	var title string
	err := db.QueryRowContext(ctx, "SELECT title FROM bxd_article WHERE id=?", id).Scan(&title)
	switch {
	case err == sql.ErrNoRows:
		return errors.Wrapf(err, "no user with id %d", id)
	case err != nil:
		return errors.Wrap(err, "query error")
	}

	return nil
}

func MyselfQuery(ctx context.Context) {
	db := NewOpenDb()

	id := 0
	rows, err := db.QueryContext(ctx, "SELECT title FROM bxd_article WHERE id=?", id)
	fmt.Printf("rows: %+v, err: %+v\n", rows, err)
	if err != nil {
		fmt.Printf("MyselfQuery: %+v\n", err)
	}
	defer rows.Close()

	titles := make([]string, 0)

	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			fmt.Printf("MyselfQuery.Next: %+v\n", err)
		}
		titles = append(titles, title)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		fmt.Printf("MyselfQuery.Close: %+v\n", rerr)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		fmt.Printf("MyselfQuery.err: %+v\n", err)
	}
	fmt.Printf("ret: %+v\n", titles)
	fmt.Printf("%s are %d years old\n", strings.Join(titles, ", "), id)
}
