package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getInfo() string {
	return fmt.Sprintf("%s is of age %d", s.name, s.age)
}

func (s *Student) setAge(newAge int) {
	s.age = newAge
}

func (s Student) getAvgGrade() float32 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	maxGrade := 0
	for _, grade := range s.grades {
		if maxGrade < grade {
			maxGrade = grade
		}
	}
	return maxGrade
}

func main() {
	std1 := Student{"James", []int{12, 35, 29, 39, 95, 21}, 19}
	fmt.Println(std1.getInfo())
	std1.setAge(100)
	fmt.Println(std1.getInfo())
	fmt.Println(std1.getAvgGrade())
	fmt.Println(std1.getMaxGrade())
}
