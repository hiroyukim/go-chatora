package main

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type KeyValueDB struct {
	DB        *sql.DB
	TableName string
}

func newKeyValueDB(config *Config) *KeyValueDB {
	db, err := sql.Open(config.getDriverName(), config.getDataSourceName())

	if err != nil {
		panic(err)
	}

	return &KeyValueDB{db, config.getTableName()}
}

func (kvdb *KeyValueDB) getValue(key string) (string, error) {
	var value string
	//TODO multi value
	//TODO QueryRowContext
	err := kvdb.DB.QueryRow("SELECT `value` FROM `"+kvdb.TableName+"` WHERE `key`=?", key).Scan(&value)
	switch {
	case err == sql.ErrNoRows:
		return value, errors.New("No value with that key.")
	case err != nil:
		return value, err
	default:
		return value, nil
	}
}

func (kvdb *KeyValueDB) close() {
	kvdb.DB.Close()
}
