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
	UDACITY_COURSES = "https://www.udacity.com/public-api/v0/courses"
	// COURSERA_CATEGORIES = "https://api.coursera.org/api/catalog.v1/categories"
)


// type courseraCategory struct {
// 	Id int
// 	Name string
// 	Description string
// }

type udacityCourse struct {
	Key string // this is the unique identifier it seems

	Title string
	Subtitle string

	Syllabus string
	Summary string

	Instructors []struct {
		Name string
	}
}


type udacityCoursesWrapper struct {
	Courses []udacityCourse
}




func udacityGetCourses() ([]udacityCourse, error) {
	body, err := Get(UDACITY_COURSES)
	if err != nil { return nil, err }

	var wrapper udacityCoursesWrapper
	err = json.Unmarshal(body, &wrapper)
	if err != nil { return nil, err }

	return wrapper.Courses, nil
}

// func courseraGetCategories() ([]courseraCategory, error) {
// 	body, err := Get(COURSERA_CATEGORIES)
// 	if err != nil { return nil, err }

// 	var wrapper courseraCategoriesWrapper
// 	err = json.Unmarshal(body, &wrapper)
// 	if err != nil { return nil, err }

// 	return wrapper.Elements, nil
// }


// func courseraFillCategories(courses []courseraCourse, categories []courseraCategory) {
// 	// convert categories to map
// 	categoryMap := make(map[int]string)
// 	for _, c := range categories {
// 		categoryMap[c.Id] = c.Name
// 	}


// 	// now go through courses and fill CategoriesString
// 	for courseIndex, course := range courses {
// 		categoriesString := make([]string, len(course.Links.Categories))
// 		for catIndex, catId := range course.Links.Categories {
// 			categoriesString[catIndex] = categoryMap[catId]
// 		}
// 		// set it
// 		courses[courseIndex].CategoriesString = categoriesString
// 	}
// }