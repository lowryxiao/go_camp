package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	perrors "github.com/pkg/errors"
)

func main() {
	if err := getData(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("get data from db successfully")

}

func getData() error {
	var data string
	db, err := sql.Open("godror", "user/pass@host:port/sid")
	if err != nil {
		return perrors.Wrap(err, "open db failed")
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "sql query")
	if err != nil {
		return perrors.WithMessage(err, "execute sql failed")
	}

	for rows.Next() {
		if err := rows.Scan(&data); err != nil {
			return perrors.WithMessage(err, "ErrNoRows")
		}

	}
	return nil
}
