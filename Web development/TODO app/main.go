package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Models aka Database - file
type Course struct {
	Name     string  `json:"coursename"`
	CourseID string  `json:"courseID"`
	Author   *author `json:"author"`
	Password string  `json:"-"`
	Price    int     `json:"price"`
}

type author struct {
	Name string `json:"AuthorName"`
	Age  int    `json:"age"`
}

//controllers aka handlers - file
func getAllcourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html;charset= utf-8")
	w.Write([]byte("<h1>All the course</h1>"))
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-type", "application/json")
	if r.Method == "GET" {
		val := r.URL.Query().Get("courseID")
		for _, v := range courses {
			if val == v.CourseID {
				json.NewEncoder(w).Encode(v)
			}
		}

	}
	//json.NewEncoder(w).Encode(co) : single line expression.
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding a New course to DB")
	w.Header().Set("Content-type", "application/json")
	//if body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode("JSON Body is nil")
		return
	}

	// if body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(course)

	//generatig UUID
	rand.Seed(time.Now().Unix())
	course.CourseID = strconv.Itoa(rand.Intn(1000))

	//Appending to database
	Courses = append(Courses, course)

}

// dummy database
var Courses []Course

func main() {
	http.Handle("/", http.HandlerFunc(getAllcourses)) //type-castiing
	http.HandleFunc("/:courseID", getCourse)
	http.HandleFunc("/addCourse", addCourse)

	fmt.Println("Starting Server ..9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
