package database


import (
	"testing"
)

func TestDatabaseConnect(t *testing.T) {
	_, err := DatabaseConnect()

	if err != nil { t.Error(err) }
}