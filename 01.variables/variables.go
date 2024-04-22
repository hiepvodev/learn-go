package main

import "fmt"

func GetMoney() (int, int) {
	return 100, 150
}

func main() {
	var userName string = "anonystick"

	println("userName: ", userName)

	fmt.Printf("Type of user: %T\n", userName)

	// 2. short name
	age := 40
	fmt.Println("age: ", age)

	//error:
	var email = "anonystick@gmail.com"
	// email := "anonystick@gmail.com"

	fmt.Println("email: ", email)

	// var a = 1
	// var b = 2
	// // var tmp int

	// // tmp = a
	// // a = b
	// // b = tmp

	// b, a = a, b
	// fmt.Println(a, b)

	m1, _ := GetMoney()
	_, m2 := GetMoney()

	fmt.Println("m1: ", m1)
	fmt.Println("m2: ", m2)

	var a = 40
	var b = 50

	c := a + b
	fmt.Println("c: ", c)

	var PI int = 4
	fmt.Println("PI: ", PI)
}