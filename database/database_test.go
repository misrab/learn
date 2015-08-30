package database


import (
	"testing"

	"os"
)

func TestDatabaseConnect(t *testing.T) {
	url := os.Getenv("POSTGRESQL_LOCAL_URL")
	session, err := DatabaseConnect(url)
	defer session.Close()

	if err != nil { t.Error(err) }
}