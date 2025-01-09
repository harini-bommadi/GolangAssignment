package assignments

import "fmt"

type Student struct {
	Name  string
	Marks []int
}

func (s *Student) AddMark(newMark int) {
	s.Marks = append(s.Marks, newMark)
	fmt.Println("Updated Marks: ", s.Marks)
}

func (s *Student) CalculateAverage() float64 {
	sum := 0
	for i := 0; i < len(s.Marks); i++ {
		sum += s.Marks[i]
	}
	if len(s.Marks) == 0 {
		return 0
	} else {
		return float64(sum) / float64(len(s.Marks))
	}
}

func Stud() {
	var s1 Student
	fmt.Println("Enter Name of Student: ")
	fmt.Scan(&s1.Name)
	var size int
	fmt.Print("Please enter the number of subjects for which you want to provide marks: ")
	fmt.Scan(&size)
	s1.Marks = make([]int, size)
	fmt.Println("Enter Marks")
	for i := 0; i < size; i++ {
		fmt.Scan(&s1.Marks[i])
	}
	fmt.Printf("Name: %s", s1.Name)
	fmt.Println()
	fmt.Print("Marks: ")
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", s1.Marks[i])
	}
	var newMark int
	fmt.Println()
	fmt.Print("Add newMarks: ")
	fmt.Scan(&newMark)
	s1.AddMark(newMark)
	fmt.Println("calculated Average of Marks: ", s1.CalculateAverage())
}
