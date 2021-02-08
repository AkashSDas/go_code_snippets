package main

import "fmt"

func main() {
	// fmt.Println(add(10, 20))

	// fmt.Println(greeting("James"))

	// sumResult, subResult := addAndSub(10, 20)
	// fmt.Println(sumResult, subResult)

	// fmt.Println(doubleMe(1, 2))

	// fmt.Println(database("stupid"))

	// myFunc := whoTheFunc
	// myFunc("👩")

	// anonymousFunc := func(drink string) string {
	// 	return fmt.Sprintf("Where's my %s", drink)
	// }
	// fmt.Println(anonymousFunc("🍹"))

	paisaDouble := outerFunc(display)
	fmt.Println(paisaDouble(10))
}

func display(emoji string) string {
	return fmt.Sprintf("Son of a gun: %s", emoji)
}

func outerFunc(display func(emoji string) string) func(moneyAmount int) int {
	info := display("👨")
	fmt.Println(info)

	return func(moneyAmount int) int {
		return moneyAmount * 2
	}
}

func whoTheFunc(emoji string) {
	fmt.Printf("I am a %s\n", emoji)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 + num2
}

func addAndSub(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func doubleMe(num1, num2 int) (num1Double, num2Double int) {
	num1Double = num1 * 2
	num2Double = num2 * 2
	return
}

func checkPermission(permission bool) {
	// fmt.Println(permission)
	if permission {
		fmt.Println("I smell success 🤑")
	} else {
		fmt.Println("Who the fuck do you think you are 🖕")
	}
}

func database(password string) (allowAccess bool) {
	defer checkPermission(allowAccess)

	if password == "stupid" {
		allowAccess = true
	} else {
		allowAccess = false
	}
	fmt.Printf("Permission given: %t\n", allowAccess)
	return
}

func greeting(name string) string {
	return fmt.Sprintf("%s hello there", name)
}
