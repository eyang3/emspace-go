package repositories

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	// this is the underlying driver
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func connect() *sqlx.DB {
	viper.SetConfigFile("./conf/db.json")
	err := viper.ReadInConfig()
	hostname := viper.GetString("host")
	username := viper.GetString("username")
	database := viper.GetString("database")
	port := viper.GetString("port")
	password := os.Getenv("POSTGRES_PW")
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, hostname, port, database)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		return nil
	}
	return db

}

func close() {
	DB.Close()
}

func String2NullString(value string) sql.NullString {
	m := sql.NullString{}
	m.String = value
	m.Valid = true
	return m
}

func Int642NullInt64(value int64) sql.NullInt64 {
	m := sql.NullInt64{}
	m.Int64 = value
	m.Valid = true
	return m
}

var DB = connect()
