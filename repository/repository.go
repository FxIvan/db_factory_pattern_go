package repository

import (
	"fmt"

	"github.com/fxivan/db_go_abstract_factory/configuration"
	mongodb "github.com/fxivan/db_go_abstract_factory/repository/mogodb"
	"github.com/fxivan/db_go_abstract_factory/repository/mysql"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
}

func New(config *configuration.Configuration) (Repository, error) {
	var repo Repository
	var err error

	switch config.Engine {
	case "mysql":
		return mysql.New(config)
	case "mongodb":
		repo = mongodb.New()
	default:
		err = fmt.Errorf("invalid engine %s", config.Engine)
	}

	return repo, err
}
