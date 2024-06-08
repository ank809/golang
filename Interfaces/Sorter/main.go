package main

import "fmt"

type Sorter interface {
	Sort([]int) []int
}

type AscendingSorter struct{}
type DescendingSorter struct{}

func (a AscendingSorter) Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
func (d DescendingSorter) Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] < arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func SortArray(s Sorter, arr []int) []int {
	return s.Sort(arr)
}

func main() {
	a := AscendingSorter{}
	arr := []int{3, 2, 5, 4, 8, 1}
	fmt.Println(SortArray(a, arr))
	d := DescendingSorter{}
	fmt.Println(SortArray(d, arr))
}
