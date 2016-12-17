package main

import (
	"fmt"
	"github.com/xujianhuaap/gostudy/exerse"
)

var s *exerse.Student
func main()  {
	fmt.Printf("%v",s)
}
func init() {
	s=&exerse.Student{Name:"li",Age:28,Gender:false}
}