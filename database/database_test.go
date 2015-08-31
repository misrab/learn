package database


import (
	"testing"

	"os"
)

func TestDatabaseConnect(t *testing.T) {
	url := os.Getenv("POSTGRESQL_LOCAL_URL")
	session, err := DatabaseConnect(url)
	if err != nil { t.Error(err) }
	defer session.Close()
}