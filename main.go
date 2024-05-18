package main

import (
	"github.com/fxivan/db_go_abstract_factory/configuration"
	"github.com/fxivan/db_go_abstract_factory/repository"
)

func main() {

	config := configuration.Configuration{
		Engine:   "mysql",
		Host:     "localhost",
		Port:     33060,
		User:     "root",
		Password: "secret",
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
