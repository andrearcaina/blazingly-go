package database

import (
	"database/sql"
	"fmt"
	"github.com/blazingly-go/crud-api/models"
)

func GetAll[T models.ModelFields](db *sql.DB, tableName string, prototype T) ([]T, error) {
	query := fmt.Sprintf("SELECT * FROM"+" %s", tableName)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []T
	for rows.Next() {
		item := prototype.NewInstance()

		err = rows.Scan(item.GetFields()...) // spread operator (...) to unpack the fields

		if err != nil {
			return nil, err
		}

		list = append(list, item.(T))
	}

	return list, nil
}

func GetID[T models.ModelFields](db *sql.DB, tableName string, prototype T, id int) (T, error) {
	query := fmt.Sprintf("SELECT * FROM"+" %s WHERE id=%d", tableName, id)
	row := db.QueryRow(query)

	item := prototype.NewInstance()

	err := row.Scan(item.GetFields()...)

	if err != nil {
		return item.(T), err
	}

	return item.(T), nil
}
