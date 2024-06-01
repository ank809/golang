package main

type Student struct {
	name  string
	marks []int
	age   int
}

func (s Student) getAverage() int {
	sum := 0
	for i := 0; i < len(s.marks); i++ {
		sum += s.marks[i]
	}
	avg := sum / len(s.marks)
	return avg
}

func (s1 Student) differnce_marks(s2 Student) []int {
	x := make([]int, len(s1.marks))
	for i := 0; i < len(s1.marks); i++ {
		x[i] = s1.marks[i] - s2.marks[i]
	}
	return x
}
