package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	b := `[{"First":"james","Last":"bond","Age":45},{"First":"money","Last":"penny","Age":45}]`
	bs := []byte(b) // explicit type conversion - byte is alias for uint8

	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", bs)

	// people := []Person{} //increase readibility
	var people []Person

	// json.Unmarshal() //Unmarshal(data []byte, v interface{}) error
	if err := json.Unmarshal(bs, &people); err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
