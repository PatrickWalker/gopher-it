package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

func (sg StudentGrade) String() string {
	return fmt.Sprintf("Name : %v - Grade : %v", sg.Name, sg.Grade)
}

func main() {
	var subject string
	//english as a default
	flag.StringVar(&subject, "subject", "english", "the subject to grade")
	flag.Parse()
	grades, err := loadScoresForSubject(subject)
	if err != nil {
		fmt.Printf("Unable to grade subject %v : %v \n", subject, err)
		os.Exit(-1)
	}
	fmt.Println(grades)
}

//loadScoresForSubject
func loadScoresForSubject(subjectName string) (grades []StudentGrade, err error) {
	//get the filename
	f, err := openSubjectFile(subjectName)
	if err != nil {
		return
	}
	//closes the file at the end of this to help with resources
	defer f.Close()
	reader := csv.NewReader(f)
	//convert score to grade and add to result set
	for {
		//read a new line
		record, err := reader.Read()
		//this signifies the end of the file
		if err == io.EOF {
			//exits the for loop
			break
		}
		//this is another error
		if err != nil {
			//you could exit here but we'll just log it out and continue
			log.Println("Unable to read line", record, err)
			//continue jumps to the next itration of the for loop advancing the reader
			continue
		}
		var grade string
		//get the score converted to an int as all values are strings from csv
		score, err := strconv.Atoi(record[ScoreIndex])
		//this means we couldn't convert to an int
		if err != nil {
			log.Println("Error parsing score for record : ", record, err)
			grade = "U"
		} else {
			//go get the grade
			grade = gradeScore(score)
		}
		//append is not massively efficient as a new slice is created each time but it works
		grades = append(grades, StudentGrade{
			Student: Student{
				Name:   record[NameIndex],
				Number: record[NumberIndex],
			},
			Grade: grade,
		})
	}
	return
}

//uses the sibling path. Better if the files were in a good known absolute folder
func getFilePath(subjectName string) (filePath string) {
	return fmt.Sprintf("../files/%v.csv", subjectName)
}

//gets the file path and then returns the opened filed
func openSubjectFile(subjectName string) (reader *os.File, err error) {
	path := getFilePath(subjectName)
	//open it and read each line
	reader, err = os.Open(path)
	return
}

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
