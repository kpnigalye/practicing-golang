package main

import "fmt"

// num1 := 4234;			// not allowed outside the function

// variable with a first letter as capital is treated as public
const LoginToken string = "sdkanskdjabs"

func main()  {
	var username string = "Krishna"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal bool = false
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	//////////////////// default values //////////////////////

	var num int
	fmt.Println(num)
	fmt.Printf("Variable is of type: %T \n", num)

	var str string
	fmt.Println(str)
	fmt.Printf("Variable is of type: %T \n", str)

	var flag bool
	fmt.Println(flag)
	fmt.Printf("Variable is of type: %T \n", flag)

	//////////////////// implicit style //////////////////////
	var str1 = "sample";
	fmt.Println(str1)
	fmt.Printf("Variable is of type: %T \n", str1)

	//////////////////// no var style //////////////////////
	num1 := 4234;
	fmt.Println(num1)
	fmt.Printf("Variable is of type: %T \n", num1)

	//////////////////// accessing public variable //////////////////////
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}