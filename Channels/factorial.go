package main

func factorial(n int, c chan int) {
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	c <- fact
}

// func main(){
//    n:=5
//    a :=make(chan int)
//    go factorial(n, a)
//    fmt.Println(<-a)
//    var str string
//     fmt.Scanln(&str)
// }
