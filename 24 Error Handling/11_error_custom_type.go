package main

import (
	"fmt"
	"log"
)

type userDefinedErr struct {
	lat, long string
	err       error
}

func (u userDefinedErr) Error() string {
	return fmt.Sprintf("Vaibhav defined math error occured: %v %v %v", u.lat, u.long, u.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		ude := fmt.Errorf("user defined: square root of negative number")
		return 0, userDefinedErr{"54.3214535", "23.463667", ude}
	}
	return 10, nil
}
