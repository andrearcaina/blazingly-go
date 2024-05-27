package database

import (
	"blazingly-go/crud-api/models"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func GetAll[T models.DatabaseObject](db *sql.DB, tableName string, prototype T) ([]T, error) {
	query := fmt.Sprintf("SELECT * FROM"+" %s", tableName)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []T
	for rows.Next() {
		item := prototype.NewObject()

		err = rows.Scan(item.GetFields()...) // spread operator (...) to unpack the fields

		if err != nil {
			return nil, err
		}

		list = append(list, item.(T))
	}

	return list, nil
}

func GetID[T models.DatabaseObject](db *sql.DB, tableName string, prototype T, id int) (T, error) {
	query := fmt.Sprintf("SELECT * FROM"+" %s WHERE id=%d", tableName, id)
	row := db.QueryRow(query)

	item := prototype.NewObject()

	err := row.Scan(item.GetFields()...)

	if err != nil {
		return item.(T), err
	}

	return item.(T), nil
}

func CreateObj[T models.DatabaseObject](db *sql.DB, tableName string, prototype T) (T, error) {
	v := reflect.ValueOf(prototype).Elem()
	typeOfV := v.Type()

	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 1; i < v.NumField(); i++ {
		columns = append(columns, typeOfV.Field(i).Tag.Get("json"))
		placeholders = append(placeholders, "?")
		values = append(values, v.Field(i).Interface())
	}

	log.Println(columns, placeholders, values)

	query := fmt.Sprintf(
		"INSERT INTO"+" %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	res, err := db.Exec(query, values...)

	if err != nil {
		return prototype, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return prototype, err
	}

	v.Field(0).SetInt(id)

	return prototype, nil
}
