package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	exam1       [2]float64
	exam2       [2]float64
	assignment1 [2]float64
	assignment2 [2]float64
	grade       float64
	isPassed    bool
}
type Course struct {
	class   []Student
	average float64
}

func main() {
	Start()
}
func Start() {
	var input string
	var course Course
	var numberOfStudent int
	var student Student

	for {
		fmt.Println("Enter The Number Of Students In The Course : (type exit to quit the program)")
		fmt.Scanln(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				numberOfStudent = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for i := 0; i < numberOfStudent; i++ {
		student = readStudentGrades(i + 1)
		course.class = append(course.class, student)
	}
	course.average = calculateCourseAverage(course)
	fmt.Printf("\nThe course average is %v", course.average)

}

func calculateCourseAverage(course Course) float64 {
	sum := 0
	for _, student := range course.class {
		sum += int(student.grade)
	}
	course.average = (float64(sum) / float64(len(course.class)))
	return course.average
}

func readStudentGrades(studentNumber int) Student {
	var student Student
	var input string

	for { //Exam 1 Grade
		fmt.Println("Student", studentNumber, "- Enter the grade between 0-100 for exam 1 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.exam1[0] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}
	for { //Exam 1 Weight
		fmt.Println("Student", studentNumber, "- Enter the weight between 0-100 for exam 1 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.exam1[1] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for { //Exam 2 Grade
		fmt.Println("Student", studentNumber, "- Enter the grade between 0-100 for exam 2 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.exam2[0] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}
	for { //Exam 2 Weight
		fmt.Println("Student", studentNumber, "- Enter the weight between 0-100 for exam 2 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.exam2[1] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for { //Assignment 1 Grade
		fmt.Println("Student", studentNumber, "- Enter the grade between 0-100 for assignment 1 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.assignment1[0] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}
	for { //Assignment 1 Weight
		fmt.Println("Student", studentNumber, "- Enter the weight between 0-100 for assignment 1 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.assignment1[1] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for { //Assignment 2 Grade
		fmt.Println("Student", studentNumber, "- Enter the grade between 0-100 for assignment 2 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.assignment2[0] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}
	for { //Assignment 2 Weight
		fmt.Println("Student", studentNumber, "- Enter the weight between 0-100 for assignment 2 : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float <= 100 {
					student.assignment2[1] = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	student.grade = getStudentCalculatedGrades(student)
	student.isPassed = isStudentPassed(student)

	fmt.Println("The Student's Average:", student.grade)
	fmt.Println("Has Student Passed?", student.isPassed)

	return student
}

func getStudentCalculatedGrades(student Student) float64 {
	var grade float64
	grade += (student.exam1[0] * student.exam1[1])
	grade += (student.exam2[0] * student.exam2[1])
	grade += (student.assignment1[0] * student.assignment1[1])
	grade += (student.assignment2[0] * student.assignment2[1])
	grade = (grade / 100)

	return grade
}

func isStudentPassed(student Student) bool {
	if student.grade >= 60 {
		return true
	}
	return false
}
