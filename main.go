package main

import (
	"fmt"

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
	/*	MongoDB
		Engine:   "mongodb",
		Host:     "localhost",
		Port:     27017,
		User:     "almendra",
		Password: "1ASWWWaeq",
		DBName:   "product",
	*/

	config := configuration.Configuration{
		Engine:   "redis",
		Host:     "localhost",
		Port:     6379,
		User:     "",
		Password: "",
		DBName:   "",
	}

	repo, err := repository.New(&config)
	if err != nil {
		panic(err)
	}

	/*err = repo.Save("Arroz")
	if err != nil {
		panic(err)
	}*/

	data := repo.FindKey("comida")
	fmt.Println(data)

}
