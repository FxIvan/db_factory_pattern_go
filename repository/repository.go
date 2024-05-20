package repository

import (
	"fmt"

	"github.com/fxivan/db_go_abstract_factory/configuration"
	mongodb "github.com/fxivan/db_go_abstract_factory/repository/mogodb"
	"github.com/fxivan/db_go_abstract_factory/repository/mysql"
	"github.com/fxivan/db_go_abstract_factory/repository/redis"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
	FindKey(key string) string
}

func New(config *configuration.Configuration) (Repository, error) {
	var repo Repository
	var err error

	switch config.Engine {
	case "mysql":
		return mysql.New(config)
	case "mongodb":
		return mongodb.New(config)
	case "redis":
		return redis.New(config)
	default:
		err = fmt.Errorf("invalid engine %s", config.Engine)
	}

	return repo, err
}
