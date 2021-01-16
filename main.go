// json encoding & decoding sample
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

//Employee is employee data
type Employee struct {
	ID        int
	Title     string
	Completed bool
	CreateAt  time.Time
}

func main() {

	// data for json
	data := []Employee{
		{
			1,
			"Play with Gopher. └(●^□^●)┘",
			false,
			time.Now(),
		},
		{
			2,
			"Learn with Gopher. (￣⊆￣)",
			false,
			time.Now(),
		},
		{
			3,
			"Sleep with Gopher. (_ _).。oＯ",
			false,
			time.Now(),
		},
	}

	// Encoding to json
	file, err := json.MarshalIndent(data, "", "")
	//file, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	// json file write
	_ = ioutil.WriteFile("test.json", file, 0644)

	// json file read
	content, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("error:", err)
	}

	// undecode data output
	fmt.Println("-1-----------")
	fmt.Printf("File contents: %s\n", content)

	//for decode data
	var m []Employee

	//Decoding from json ( to m )
	if err := json.Unmarshal(content, &m); err != nil {
		fmt.Println("error:", err)
	}

	// decode data output
	fmt.Println("-2-----------")
	fmt.Println(m)
	fmt.Println("-3-----------")
	fmt.Printf("%+v\n", m)

}
