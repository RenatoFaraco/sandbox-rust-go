package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Age    int
	Grades []float32
	Status string
}

func (s *Student) UpdateStatus() {
	var sum float32 = 0
	for _, grade := range s.Grades {
		sum += grade
	}
	average := sum / float32(len(s.Grades))
	if average >= 60 {
		s.Status = "Aprovado"
	} else {
		s.Status = "Reprovado"
	}

}

func main() {
	student1 := Student{
		Name:   "João",
		Age:    20,
		Grades: []float32{70, 80, 90},
	}

	student2 := Student{
		Name:   "Maria",
		Age:    22,
		Grades: []float32{50, 60, 10},
	}

	student1.UpdateStatus()
	student2.UpdateStatus()

	fmt.Println("Nome:", student1.Name)
	fmt.Println("Idade:", student1.Age)
	fmt.Println("Notas:", student1.Grades)
	fmt.Println("Status:", student1.Status)

	fmt.Println("Nome:", student2.Name)
	fmt.Println("Idade:", student2.Age)
	fmt.Println("Notas:", student2.Grades)
	fmt.Println("Status:", student2.Status)
}
