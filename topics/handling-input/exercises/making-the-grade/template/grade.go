package main

import (
	"fmt"
)

//These constants are so if we add another field we just update the index in 1 place
const (
	NameIndex   int = 0
	NumberIndex int = 1
	ScoreIndex  int = 2
)

//Student models the info we have on a student
type Student struct {
	Name   string
	Number string
}

//StudentGrade is a structe which has the student details and grade
type StudentGrade struct {
	Student
	Grade string
}

//String implements the stringer method which means when we try and fmt.Println we get this output
func (sg StudentGrade) String() string {
	return fmt.Sprintf("Name : %v - Grade : %v", sg.Name, sg.Grade)
}

func main() {

	//setup your flag to handle the subject and pick a healthy default
	//flag.Parse()

	//use that subject to work out the file path

	//as it's a sibling directiory we can use ../file/<subject>.csv
	//fmt.Sprintf is handy for string interpolation

	//open file iterate on it and parse each record

	//you can use sub functions or whatever you like

}

//gradescore takes an integer and returns a grade. if the number is <0 or >100 we return a U
func gradeScore(score int) (grade string) {
	switch {
	case score >= 0 && score <= 49:
		grade = "F"
	case score >= 50 && score <= 59:
		grade = "D"
	case score >= 60 && score <= 69:
		grade = "C"
	case score >= 70 && score <= 79:
		grade = "B"
	case score >= 80 && score <= 89:
		grade = "A"
	case score >= 90 && score <= 100:
		grade = "A+"
	default:
		grade = "U"
	}
	return grade
}
