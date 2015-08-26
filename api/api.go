package api


import (
	// "fmt"

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


// generic unmarshal
// func Unmarshal(b []byte, wrapper interface{}) {
// 	return json.Unmarshal(b, wrapper)
// 	if err != nil { return err }

// 	re
// }