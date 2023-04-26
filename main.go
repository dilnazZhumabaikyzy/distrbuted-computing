package main

import (
	"fmt"	
    "log"
	"github.com/gomodule/redigo/redis"
)

type Course  struct{
	code string
	grade string
}

type Student struct{
	id string
	name string
	courses []Course
}
func (s Student) viewCourses(conn redis.Conn, id string ){
	_, err := redis.StringMap(conn.Do(
	   "HMGETALL",
	   id,
	))
	checkError(err)
	fmt.Println("Courses of " + id)
	// for k, v := rangeok values {
	// 	fmt.Println("Ok")
	// }
}
func (s Student) addCourse(conn redis.Conn, course Course, student Student){
    student.courses.append()
     _, err := conn.Do(
		"HMSET",
		student.id,
        "courses",
         .code,
	 )
     checkError(err)
     fmt.Println("You added the course!")
}
func (s Student) delCourse(conn redis.Conn, course Course, id string){
		_, err := conn.Do(
		"HDEL",
		id,
		"course",
	    course.code,
	)
	checkError(err)
     fmt.Println("You deleted the course!")
}
func (s Student) updateCourse(conn redis.Conn, course Course, id string){
	_, err := conn.Do(
	"HDEL",
	id,
	"course",
	course.code,
)
checkError(err)
 fmt.Println("You updated the course!")
}
func NewStudent(id int, name string) * Student{
	return &Student{
     courses: make([]Course, 7),
	 name: name,
	 id: id,
	}
}


func checkError(err error){
	if(err != nil){
	 log.Fatal(err)
	}
}
func main(){
	// Create a new Redis client
    fmt.Printf(" i am main go")
	conn, err := redis.Dial("tcp", "localhost:6379")	
    checkError(err)

	conn.Close()


}
// func CreateStudent(string name){
// 	_, err = conn.Do("HMSET", "ID", "Name", name, "Courses", )
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
