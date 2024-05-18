package main

import (
	"github.com/fxivan/db_go_abstract_factory/configuration"
	"github.com/fxivan/db_go_abstract_factory/repository"
)

func main() {
	/*	MySQL
		Engine:   "mysql",
		Host:     "localhost",
		Port:     33060,
		User:     "root",
		Password: "secret",
		DBName:   "product",
	*/

	config := configuration.Configuration{
		Engine:   "mongodb",
		Host:     "localhost",
		Port:     27017,
		User:     "almendra",
		Password: "1ASWWWaeq",
		DBName:   "product",
	}

	repo, err := repository.New(&config)
	if err != nil {
		panic(err)
	}

	err = repo.Save("Arroz")
	if err != nil {
		panic(err)
	}

	/*data := repo.Find(1)
	fmt.Println(data)*/

}
