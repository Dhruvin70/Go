package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Auther      *Auther `json:"auther"`
}

type Auther struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake database

var courses []Course

//Middleware , helper -file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	router := mux.NewRouter()

	// seeding  of the data
	courses = append(courses, Course{CourseId: "2", CourseName: "golang", CoursePrice: 599, Auther: &Auther{Fullname: "Dhruvin Gadhiya", Website: "google.com"}})
	courses = append(courses, Course{CourseId: "1", CourseName: "java", CoursePrice: 499, Auther: &Auther{Fullname: "Dhruvin patel", Website: "yahoo.com"}})

	//routing
	// get :== Retrieves data from the server.
	// post:==  It submits the processed data to a specified resource.
	router.HandleFunc("/", servehome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/create", createOneCourse).Methods("POST")
	router.HandleFunc("/update/{id}", uptadeOneCourse).Methods("PUT")
	router.HandleFunc("/delete", deleteOneCoure).Methods("DELETE")

	// listening to a port
	log.Fatal(http.ListenAndServe(":4000", router))
}

// controllers -file

// serve home route

func servehome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> welcome </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// taking help from request
	// grab id from user
	// check in arry if exists
	w.Write([]byte("<h1> getOneCourse </h1>"))
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)
	println(params)

	// loop through courses, finding matching od and return the responce

	for _, course := range courses {
		if course.CourseId == params["id"] {
			// we can encode whatever we want
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data")
		return
	}

	// TODO: check only if title is duplicate
	// loop through courses.Course.name
	for _, coursename := range courses {
		if coursename.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Already exists")
			return
		}
	}

	//generate unique id from string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func uptadeOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req

	params := mux.Vars(r)

	// loop through the value
	// remove when id found
	// add value with my id which is getting from params

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return
	// TODO : send a responce when id is not found
}

func deleteOneCoure(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	params := mux.Vars(r)

	// loop ultill id reached
	// remove from that index

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
