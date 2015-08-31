/*
	coursera api: https://tech.coursera.org/app-platform/catalog/

	it seems a course can be reached at 
	https://www.coursera.org/course/[shortname]
*/

package fetcher

import (
	// "fmt"

	"encoding/json"
)


const (
	COURSERA_COURSES = "https://api.coursera.org/api/catalog.v1/courses?includes=categories&fields=name,shortDescription,aboutTheCourse,courseSyllabus,suggestedReading"
	COURSERA_CATEGORIES = "https://api.coursera.org/api/catalog.v1/categories"
)


type courseraCategory struct {
	Id int
	Name string
	Description string
}

type courseraCourse struct {
	Id int
	Name string
	ShortDescription string

	PreviewLink string
	
	AboutTheCourse string
	CouseSyllabus string
	SuggestedReadings string
	Instructor string // TODO this is likely an id
	RecommendedBackground string

	Links struct {
		Categories []int
	}

	// my own
	CategoriesString []string
}


type courseraCoursesWrapper struct {
	Elements []courseraCourse
}
type courseraCategoriesWrapper struct {
	Elements []courseraCategory
}



func courseraGetCourses() ([]courseraCourse, error) {
	body, err := Get(COURSERA_COURSES)
	if err != nil { return nil, err }

	var wrapper courseraCoursesWrapper
	err = json.Unmarshal(body, &wrapper)
	if err != nil { return nil, err }

	return wrapper.Elements, nil
}

func courseraGetCategories() ([]courseraCategory, error) {
	body, err := Get(COURSERA_CATEGORIES)
	if err != nil { return nil, err }

	var wrapper courseraCategoriesWrapper
	err = json.Unmarshal(body, &wrapper)
	if err != nil { return nil, err }

	return wrapper.Elements, nil
}


func courseraFillCategories(courses []courseraCourse, categories []courseraCategory) {
	// convert categories to map
	categoryMap := make(map[int]string)
	for _, c := range categories {
		categoryMap[c.Id] = c.Name
	}


	// now go through courses and fill CategoriesString
	for courseIndex, course := range courses {
		categoriesString := make([]string, len(course.Links.Categories))
		for catIndex, catId := range course.Links.Categories {
			categoriesString[catIndex] = categoryMap[catId]
		}
		// set it
		courses[courseIndex].CategoriesString = categoriesString
	}
}