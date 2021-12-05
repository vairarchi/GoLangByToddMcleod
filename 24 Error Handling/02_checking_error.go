package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string
	fmt.Print("Name: ")
	if _, err := fmt.Scan(&ans1); err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	if _, err := fmt.Scan(&ans2); err != nil {
		panic(err)
	}

	fmt.Print("Fav animal: ")
	if _, err := fmt.Scan(&ans3); err != nil {
		panic(err)
	}

	fmt.Println(ans1, ans2, ans3)
}
