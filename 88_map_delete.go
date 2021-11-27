package main

import "fmt"

func main() {
	m := map[string]int{
		"james": 58,
		"jam":   56,
		"mfs":   21,
	}
	fmt.Println(m)

	delete(m, "james")
	fmt.Println(m)

	// if deleteing a key-valye not present no error is thrown
	delete(m, "vaibhav")

	// we use comma oka idiom
	if _, ok := m["vaibhav"]; ok {
		delete(m, "vaibhav")
	} else {
		fmt.Println("Cant delete as key not present")
	}

}
