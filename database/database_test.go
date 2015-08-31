package database


import (
	"testing"

	// "fmt"
	"os"
	"errors"
)

func TestDatabaseBasic(t *testing.T) {
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
	model := Model{Name:"foo", Value:64}
	session.Insert(&model)


	// try and get it back
	records := []Model{}
	session.FindAll(&records)

	if len(records) != 1 {
		t.Error(errors.New("Couldn't retrieve the record we just inserted"))
	}

	record := records[0]
	if record.Name != model.Name || record.Value != model.Value {
		t.Error(errors.New("Record retrieved didn't have the right values"))
	}
}