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

func NewKeyValueDB(config *Config) *KeyValueDB {
	db, err := sql.Open(config.GetDriverName(), config.GetDataSourceName())

	if err != nil {
		panic(err)
	}

	return &KeyValueDB{db, config.GetTableName()}
}

func (kvdb *KeyValueDB) getValue(key string) (string, error) {
	var value string
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

func (kvdb *KeyValueDB) Close() {
	kvdb.DB.Close()
}
