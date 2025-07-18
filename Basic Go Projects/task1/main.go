package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Student struct {
	name         string
	noOfSubjects int
	subjects     map[string]float32
	average      float32
}

func (s *Student) AddSubject(subject string, score float32) {
	s.subjects[subject] = score
}

func (s *Student) CalculateAverage() float32 {
	var total float32
	for _, score := range s.subjects {
		total += score
	}
	s.average = total / float32(s.noOfSubjects)
	return s.average
}

func (s *Student) print() {
	fmt.Printf("\n=================================\n")
	fmt.Printf("   Student Name: %v\n", s.name)
	fmt.Printf("=================================\n")

	fmt.Println("┌────────────────────┬────────────┐")
	fmt.Printf("│ %-18s │ %-10s │\n", "Subject Name", "Score")
	fmt.Println("├────────────────────┼────────────┤")

	for key, value := range s.subjects {
		fmt.Printf("│ %-18s │ %-10.2f │\n", key, value)
		fmt.Println("├────────────────────┼────────────┤")
	}
	s.CalculateAverage()
	fmt.Printf("│ %-18s │ %-10.2f │\n", "Average Score", s.average)
	fmt.Println("└────────────────────┴────────────┘")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var stud Student
	stud.subjects = make(map[string]float32)

	fmt.Print("Enter Student Name: ")
	nameInput, _ := reader.ReadString('\n')
	stud.name = strings.TrimSpace(nameInput)

	fmt.Print("Enter number of subjects: ")
	countStr, _ := reader.ReadString('\n')
	countStr = strings.TrimSpace(countStr)
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		fmt.Println("Invalid number of subjects.")
		return
	}
	stud.noOfSubjects = count

	for i := 0; i < stud.noOfSubjects; i++ {
		fmt.Printf("Subject %d Name: ", i+1)
		subjName, _ := reader.ReadString('\n')
		subjName = strings.TrimSpace(subjName)

		fmt.Print("Score: ")
		scoreStr, _ := reader.ReadString('\n')
		scoreStr = strings.TrimSpace(scoreStr)
		score, err := strconv.ParseFloat(scoreStr, 32)
		if err != nil || score < 0 {
			fmt.Println("Invalid score.")
			return
		}
		stud.AddSubject(subjName, float32(score))
		clearConsole()
	}

	stud.print()
}
