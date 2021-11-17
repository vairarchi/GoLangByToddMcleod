/*
Map - introduction
A map is a key-value store.
This means that you store some value and you access that value by a key.
For instance, I might store the value “phoneNumber”
and access it be the key “friendName”.

This way I could enter my friend’s name and have the map return their phone number.
The syntax for a map is map[key]value.
The key can be of any type which allows comparison
(eg, I could use a string or an int, for example,
	as I can compare if two strings are equal, or if two ints are equal).

	It is important to note that maps are unordered.
	If you print out all of the keys and values in a map,
	they will print out in random order.
	The comma ok idiom is also covered in this video.
	A map is the perfect data structure when you need to look up data fast.

*/

package main

import "fmt"

func main() {
	// composite literal for map
	m := map[string]int{
		"james":   38,
		"giannis": 28,
	}

	fmt.Println(m)
	fmt.Println(m["james"])
	// accessing key of something not present in the map
	fmt.Println(m["Vaibhav"])

	// v, ok := m["Vaibhav"]
	// fmt.Println(v)
	// fmt.Println(ok)

	// idiomatic expression
	if v, ok := m["Vaibhav"]; ok {
		fmt.Println("If print: ", v)
	} else {
		fmt.Println("Key not prresent")
	}
}
