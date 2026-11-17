package main

import "fmt"

type Student struct {
	Name  string
	Score []float32
}

func (s *Student) average() float32 {
	cont := len(s.Score)
	var score float32 = 0
	if cont == 0 {
		return 0
	}
	for i := 0; i < cont; i++ {
		score += s.Score[i]
	}
	score = score / float32(cont)
	return score
}

func (s *Student) Grade() string {
	score := s.average()
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func (s *Student) StudentINFO(name string, point []float32) {
	s.Name = name
	s.Score = point
}

func main() {
	student := []Student{
		{Name: "Kuy", Score: []float32{30, 40, 70}},
		{Name: "Hee", Score: []float32{100, 100, 100}},
	}

	for i := 0; i < len(student); i++ {
		s := &student[i]
		fmt.Printf("%s  avg %.1f  grade %s\n", s.Name, s.average(), s.Grade())
	}
}
