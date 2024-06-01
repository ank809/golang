package main

import "fmt"

type ContactInfo struct {
	email    string
	mobileNo int
}
type Address struct {
	country string
	city    string
}

type User struct {
	firstName string
	lastName  string
	contact   ContactInfo
	address   Address
}

func UserDetails() {
	user := User{
		firstName: "Ankita",
		lastName:  "Thakur",
		contact: ContactInfo{
			email:    "xyz23@gmail.com",
			mobileNo: 1234564567890,
		},
		address: Address{
			country: "India",
			city:    "Jalandhar",
		},
	}
	fmt.Println(user)
	fmt.Printf("%+v", user)
}
