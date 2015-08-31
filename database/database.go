/*
	database/database.go

	the interface to whatever database we use to store
	courses, books etc.

	DatabaseConnect() determines the implementation under use
*/

package database


import (
	"os"

	// "database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

/*
	Interface
*/

type Database interface {
	Insert(interface{}) error


	FindAll(interface{})

	Close()
}


/*
	"Constructor"
*/

/* This determines what implementation we're using. 

 It expects pointers to the structs for each model.
 e.g. type user struct {}. Then pass &user{}
*/
func DatabaseConnect(url string, models...interface{}) (Database, error){
	db, err := gorm.Open("postgres", url)
	if err != nil { return nil, err }




	// database setup
	env := os.Getenv("ENV")
	if env == "" || env == "development" {
		db.DropTableIfExists(models...)

		// set logging
		db.LogMode(true)
	}

	db.CreateTable(models...)
		




	// return
	session := new(Session)
	session.db = &db
	return session, nil
}

// alias to extend implementation
type Session struct {
	db *gorm.DB
}

/*
	Initial implementation
*/

func (s *Session) Close() {
	s.db.Close()
}


func (s *Session) Insert(v interface{}) error {
	if err := s.db.Create(v).Error; err != nil {
  	return err
	}

	return nil
}


func (s *Session) FindAll(v interface{}) {
	s.db.Find(v)
}