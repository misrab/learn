/*
	test getting data from different sources
*/

package fetcher

import (
	"testing"
	// "fmt"
	// "time"
)


func TestCoursera(t *testing.T) {
	// get courses
	courses, err := courseraGetCourses()
	if err != nil {
		t.Error(err)
	}

	// get categories
	categories, err := courseraGetCategories()
	if err != nil {
		t.Error(err)
	}

	// fill up categories in string form for each course
	courseraFillCategories(courses, categories)

	// for _, c := range courses {
	// 	fmt.Printf("%+v\n", c)
	// }

}


func TestUdacity(t *testing.T) {
	// get courses
	_, err := udacityGetCourses()
	if err != nil {
		t.Error(err)
	}

	// for _, c := range courses {
	// 	fmt.Printf("%+v\n", c)
	// }

}



func TestUpdate(t *testing.T) {
	Update()
}