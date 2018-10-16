package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	OmitEmpty  string `json:"omit,omitempty"`
	OmitAlways string `json:"-"`
}

func main() {

	data := []byte(`   {
        "id": "k34rAT4",
        "title": "My Awesome App"}`)

	var app App
	err := json.Unmarshal(data, &app)
	if err != nil {
		fmt.Printf("Error unmarshalling data")
	}
	fmt.Println(app)

	a1 := App{
		Id:         "all fields",
		Title:      "title",
		OmitEmpty:  "here",
		OmitAlways: "yeo",
	}
	a1Json, err := json.Marshal(a1)
	fmt.Println(string(a1Json))
	if err != nil {
		fmt.Printf("Error unmarshalling data")
	}
	a2 := App{
		Id:    "all fields",
		Title: "title",
		//Empty string is empty by json
		OmitEmpty:  "",
		OmitAlways: "foo",
	}
	a2Json, err := json.Marshal(a2)

	fmt.Println(string(a2Json))
	if err != nil {
		fmt.Printf("Error unmarshalling data")
	}

}
