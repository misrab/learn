/*
	api/api.go

	generic functions for this module
*/


package fetcher


import (
	// "fmt"
	"time"

	"io/ioutil"
	"net/http"
)



// generic get
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil { return nil, err }
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// task scheduling
// e.g. database update from apis
// does said task at interval. logs things.
func SetInterval(interval time.Duration, quit chan struct{}, f func()) {
	ticker := time.NewTicker(interval)
	// quit := make(chan struct{})
	go func() {
	    for {
	       select {
	        case <- ticker.C:
	            // do stuff
	            f()
	        case <- quit:
	            ticker.Stop()
	            return
	        }
	    }
	 }()
}


/*
	This is the task we'll run intermittently.
	It goes through each source and updates the 
	appropriate database collection, by id, including new inserts.
*/

func Update() {
	updateCoursera()
}