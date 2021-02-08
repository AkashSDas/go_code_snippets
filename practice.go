package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

}

func numToNumsList(num int) (numsList []string) {
	// numStr := strconv.FormatInt(int64(num), 10)
	numStr := strconv.Itoa(num)
	for _, char := range numStr {
		numsList = append(numsList, fmt.Sprintf("%c", char))
	}
	return
}

func checkPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func takeUserInput() map[string]string {
	info := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fields := []string{"firstName", "lastName", "email", "password", "confirmPassword"}
	for _, field := range fields {
		fmt.Printf("%s: ", field)
		scanner.Scan()
		info[field] = scanner.Text()
	}
	return info
}

func passwordValid(password, confirmPassword string) bool {
	if password == confirmPassword {
		return true
	}
	return false
}

func login() {
	isLoggedIn := false

	fmt.Println("🍔 Please Login\n")
	for !isLoggedIn {
		info := takeUserInput()
		isLoggedIn = passwordValid(info["password"], info["confirmPassword"])
		if !isLoggedIn {
			fmt.Println("👊 Password invalid please try again\n")
		} else {
			fmt.Println("🍺 Cheers your logged in")
		}
	}
}

func shopping(numOfItemsToBuy int) []string {
	var list []string = []string{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🍕 Enter items")
	for i := 0; i < numOfItemsToBuy; i++ {
		fmt.Printf("%d. ", i+1)
		scanner.Scan()
		list = append(list, scanner.Text())
	}
	return list
}

func addEmojiToListItems(list []string) []string {
	var emojiList [7]string = [7]string{"🍺", "🍎", "🍔", "🌯", "🌮", "🎂", "🍶"}
	for idx, item := range list {
		list[idx] = fmt.Sprintf("%s %s", emojiList[rand.Intn(len(emojiList))], item)
	}
	return list
}

func takeCalInput() (num1, num2 float64) {
	fmt.Println("The world's 🍺 best calculator")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Num 1: ")
	scanner.Scan()
	num1, _ = strconv.ParseFloat(scanner.Text(), 32)
	fmt.Print("Num 2: ")
	scanner.Scan()
	num2, _ = strconv.ParseFloat(scanner.Text(), 32)
	return
}

func calculator(operation string) float64 {
	// Taking user input
	num1, num2 := takeCalInput()

	// Using switch cases
	switch operation {
	case "add":
		return num1 + num2
	case "sub":
		return num1 - num2
	case "multiply":
		return num1 * num2
	case "divide":
		return num1 / num2
	default:
		return -1
	}
}
