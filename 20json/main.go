package main

import (
	"encoding/json"
	"fmt"
)

// data comes in byte form as we know
// private
type course struct {
	Name		string		`json:"coursename"`
	Price		int
	Platform	string		`json:"website"`
	Password 	string		`json:"-"`
	Tags		[]string	`json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {	// create JSON data
	dpCourses := []course{	// slice of type courses
		{"ReactJS", 299, "diversepixel.com", "abc123", []string{"web-dev", "js"}},
		{"RubyOnRails", 199, "diversepixel.com", "bcd234", []string{"web-dev", "js"}},
		{"Golang", 299, "diversepixel.com", "abc723", nil},
	}

	//package as JSON data
	//interface is passed, word for struct
	// finalJson, err := json.Marshal(dpCourses)
	// finalJson, err := json.MarshalIndent(dpCourses, "diverse", "\t")
	finalJson, err := json.MarshalIndent(dpCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "RubyOnRails",
		"Price": 199,
		"website": "diversepixel.com",
		"tags": ["web-dev","js"]
        }
	`)

	var dpCourse course

	checkValid := json.Valid(jsonDataFromWeb)	// doesn't return error

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &dpCourse) // reference, we don't want copies to be passed
		fmt.Printf("%#v", dpCourse)
	} else {
		fmt.Println("THAT JSON IS NOT VALID!")
	}

	// some cases for getting key-value pairs

	// since value is not guaranteed, wrote interface
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}