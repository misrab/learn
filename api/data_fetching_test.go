/*
	test getting data from different sources
*/

package api

import (
	"testing"
	"fmt"
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

	for _, c := range courses {
		fmt.Printf("%+v\n", c)
	}

}