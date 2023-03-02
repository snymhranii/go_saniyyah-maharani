package main

import (
	"fmt"
)

type student struct {
	name  string
	score []float64
}

func (s *student) averageScore() float64 {
	var total float64
	for _, score := range s.score {
		total += score
	}
	return total / float64(len(s.score))
}

func main() {
	var students [5]student
	students[0] = student{name: "Rizky", score: []float64{80}}
	students[1] = student{name: "Kobar", score: []float64{75}}
	students[2] = student{name: "Ismail", score: []float64{70}}
	students[3] = student{name: "Umam", score: []float64{75}}
	students[4] = student{name: "Yopan", score: []float64{60}}

	var totalScore float64
	for _, s := range students {
		totalScore += s.averageScore()
	}
	average := totalScore / float64(len(students))
	fmt.Printf("Rata-rata skor siswa: %.2f\n", average)

	var minScore, maxScore float64
	var minStudent, maxStudent student
	for i, s := range students {
		if i == 0 {
			minScore = s.averageScore()
			maxScore = s.averageScore()
			minStudent = s
			maxStudent = s
		} else {
			if s.averageScore() < minScore {
				minScore = s.averageScore()
				minStudent = s
			}
			if s.averageScore() > maxScore {
				maxScore = s.averageScore()
				maxStudent = s
			}
		}
	}
	fmt.Printf("Siswa dengan skor minimum: %s (%.2f)\n", minStudent.name, minScore)
	fmt.Printf("Siswa dengan skor maksimum: %s (%.2f)\n", maxStudent.name, maxScore)
}
