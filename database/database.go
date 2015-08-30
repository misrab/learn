/*
	database/database.go

	the interface to whatever database we use to store
	courses, books etc.

	DatabaseConnect() determines the implementation under use
*/

package database


import (
	"gopkg.in/mgo.v2"
)

/*
	Interface
*/

type Database interface {

}


/*
	"Constructor"
*/

// this determines what implementation we're 
// using
func DatabaseConnect() {

	return nil, nil
}


/*
	Initial implementation
*/

