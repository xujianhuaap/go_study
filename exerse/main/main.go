package main

import (
	"fmt"
	"github.com/xujianhuaap/gostudy/exerse"
	"reflect"
)

var s *exerse.Student
func main()  {
	fmt.Printf("%v",s)
	fmt.Println()
	fmt.Printf("relect type %v",reflect.TypeOf(s))
	fmt.Println()
	fmt.Printf("relect value %v",reflect.ValueOf(s))
}
func init() {
	s=&exerse.Student{Name:"li",Age:28,Gender:false}
}