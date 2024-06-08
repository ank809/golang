package main

import "fmt"

//	type error interface{
//		Error() string
//	}

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("No msg")
	}
}

func main() {
	e := MyError{Message: "This is a custom error "}
	PrintError(e)
	PrintError(nil)

}
