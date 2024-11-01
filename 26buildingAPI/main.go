package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

//  middlewares, helper - seperate file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Building API in Go")
}

//  Controllers stay in different files

// server home route controller
//
//	writer and reader goes inside the function
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API for home</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the id from the request
	params := mux.Vars(r)
	id := params["id"]

	// loop through the courses to find the course with the matching id

	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("No course found with %s id", id))
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// check if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// half a** data ? like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	// generate unique id(string) append to course
	rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}
