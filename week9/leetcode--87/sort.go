package main 

import "sort"
import "fmt"
import "strings"
func main(){
	str := strings.Split("123", "")
	sort.Strings(str)
	str2 := strings.Join(str, "")
	fmt.Println(str2)
}