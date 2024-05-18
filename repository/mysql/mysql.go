package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/fxivan/db_go_abstract_factory/configuration"
)

type MySQL struct {
	DB *sql.DB
}

func New(config *configuration.Configuration) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mysql: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping mysql: %w", err)
	}

	return &MySQL{DB: db}, nil
}

func (m *MySQL) Find(id int) string {
	return "function Data | from mysql"
}

func (m *MySQL) Save(data string) error {
	query := "INSERT INTO my_table (data) VALUES (?)"
	_, err := m.DB.Exec(query, data)
	if err != nil {
		return fmt.Errorf("could not insert into mysql: %w", err)
	}
	fmt.Printf("Function Save | data to mysql %s", data)
	return nil
}
