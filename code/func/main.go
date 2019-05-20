package main


import (
	"fmt"
"code/func/read")



func main() {
	//type content func(string) string
	content := read.GetContent()
	fmt.Println(content("5658"))
}