package database


import (
	"testing"

	// "fmt"
	"os"
)

func TestDatabaseConnect(t *testing.T) {
	url := os.Getenv("POSTGRESQL_LOCAL_URL")

	type Model struct {
		Id int64
		Name string
		Value int64
	}

	session, err := DatabaseConnect(url, &Model{})
	if err != nil { t.Error(err) }
	defer session.Close()


	// test insert
	// model := Model{Name:"foo", Value:64}
	// session.Insert(&model)

	// records := []Model{}
	// session.FindAll(&records)
	// // fmt.Print(len(records))
	// for _, v := range records {
	// 	fmt.Printf("%+v", v)
	// }
}