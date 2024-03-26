package main

func addelement(a []int, c chan int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	c <- sum
}

// func main(){
//     a:=[]int{1, 2, 3, 4, 5}
//     c:=make(chan int)
//     go addelement(a, c)
//     fmt.Println("Sum is", <-c);
//     var str string
//     fmt.Scanln(&str)
// }
