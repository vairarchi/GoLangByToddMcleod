package main

import (
	"fmt"
	"os"

	"github.com/vairarchi/GoLangByToddMcLeod/data"
)

func main() {
	ids := data.EmployeeIDs
	attendance := data.AttendanceDetails

	fmt.Println(len(ids), len(attendance))
	length := len(ids)

	var mongoScript []string

	queryPt1 := `db.getCollection('employeeAttendance').update({ "employeeCode": "`
	queryPt2 := `"},
    {
        "$push": {`
	queryPt3 := `
}
});`
	for i := 0; i < length; i++ {
		mongoQuery := queryPt1 + ids[i] + queryPt2 + attendance[i] + queryPt3
		mongoScript = append(mongoScript, mongoQuery)
	}
	// fmt.Println(mongoScript)
	printLines("script.txt", mongoScript)
}

func printLines(filePath string, values []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}
