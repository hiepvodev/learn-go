package main

import "fmt"

func main() {
	var a int = 10
	var b int32 = 20

	fmt.Println(a + int(b))

	//float 
	var c float32 = 10.34
	fmt.Println("c", int(c))

	//module %
	var i int32
	var j int64
	i, j = 2, 3

	if i == 2 || j == 3 {
		fmt.Println("equal")
	}

	//Boolean
	ok := true
	fmt.Println("ok", ok)

	m := 1
	if m == 1 {
		fmt.Println("m is 1")
	}

	//string
	s1 := "hello"
	s2 := " tip go"
	s := `row1\\r\n
	row2`

    fmt.Println(s1, s2, s)

	s3 := s1 + s2

	fmt.Println(s3)


}