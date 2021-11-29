/*
Hands-on exercise #5
Starting with this code, sort the []user by
age
last
Also sort each []string “Sayings” for each user
print everything in a way that is pleasant
solution: https://play.golang.org/p/8RKkdtLl6w
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []user

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].Last < b[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for i, vu := range users {
		fmt.Println("User: ", i, "  ", vu.First, vu.Last, vu.Age)
		for _, v := range vu.Sayings {
			fmt.Println("\t\t", v)
		}
	}

	// your code goes here
	sort.Sort(ByAge(users))
	for i, vu := range users {
		fmt.Println("User: ", i, "  ", vu.First, vu.Last, vu.Age)
		for _, v := range vu.Sayings {
			fmt.Println("\t\t", v)
		}
	}
	sort.Sort(ByName(users))
	for i, vu := range users {
		fmt.Println("User: ", i, "  ", vu.First, vu.Last, vu.Age)
		for _, v := range vu.Sayings {
			fmt.Println("\t\t", v)
		}
	}
}
