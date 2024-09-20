package database

import (
	"errors"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/lib/pq"
)

func IsUniqueConstraintError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return true
	}
	return false
}

func IsNotNullViolation(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23502" {
		return true
	}
	return false
}

func IsForeignKeyViolation(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23503" {
		return true
	}
	return false
}

func SnakeCaseToCamelCase(inputUnderScoreStr string) string {
	parts := strings.Split(inputUnderScoreStr, "_")
	for index := range parts {
		if index != 0 {
			parts[index] = strings.Title(strings.ToLower(parts[index]))
		} else {
			parts[index] = strings.ToLower(parts[index])
		}
	}
	return strings.Join(parts, "")
}

func InitDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.Mapper = reflectx.NewMapperFunc("json", func(s string) string {
		return SnakeCaseToCamelCase(s)
	})

	return db, nil
}
