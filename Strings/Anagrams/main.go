package main

import (
	"fmt"
	"sort"
	"strings"
)

func checkAnagrams(str1, str2 string) bool {
	if (len(str1) < len(str2)) || (len(str2) < len(str1)) {
		return false
	}
	s := strings.Split(str1, "")
	s2 := strings.Split(str2, "")
	sort.Strings(s)
	sort.Strings(s2)
	fmt.Println(s, " ", s2)
	str1 = strings.Join(s, "")
	str2 = strings.Join(s2, "")
	fmt.Println(str1, " ", str2)
	return str1 == str2
}
func main() {
	str1 := "hello world"
	str2 := "world hello"

	fmt.Println(checkAnagrams(str1, str2))

}
