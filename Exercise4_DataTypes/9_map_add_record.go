/*
Hands-on exercise #9
Using the code from the previous example, add a record to your map.
Now print the map out using the “range” loop
solution: https://play.golang.org/p/_CkxAhRrDJ

*/

package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["todd_mcleod"] = []string{"Beer", "Family", "Teaching"}

	for k, v := range m {
		fmt.Println("This is record for : ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
