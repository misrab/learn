/*
	database/database.go

	the interface to whatever database we use to store
	courses, books etc.

	DatabaseConnect() determines the implementation under use
*/

package database


import (
	"database/sql"
	_ "github.com/lib/pq"
)

/*
	Interface
*/

type Database interface {
	Close()
}


/*
	"Constructor"
*/

// this determines what implementation we're 
// using
func DatabaseConnect(url string) (Database, error){
	db, err := sql.Open("postgres", url)
	if err != nil { return nil, err }

	session := new(Session)
	session.db = db
	return session, nil
}

// alias to extend implementation
type Session struct {
	db *sql.DB
}

/*
	Initial implementation
*/

func (s *Session) Close() {
	s.db.Close()
}