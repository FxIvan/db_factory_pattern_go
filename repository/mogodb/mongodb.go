package mongodb

import "fmt"

type MongoDB struct{}

func New() *MongoDB {
	return &MongoDB{}
}

func (m *MongoDB) Find(id int) string {
	return "function Data | from MongoDB"
}

func (m *MongoDB) Save(data string) error {
	fmt.Println("function Save | data to MongoDB")
	return nil
}
