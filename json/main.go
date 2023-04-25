package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//EncodingJson()
	DecodingJson()
}

type course struct {
	Name     string `json:"Cousrename"`
	Price    int
	Platform string   `json:"Delivery-Mode"`
	Password string   `json:"-"` // if you dont want to reflect your information to user using this api
	Tags     []string `json:"tags,omitempty"`
}

func EncodingJson() {
	Georgian := []course{

		{"GO", 253, "Online", "abc123", []string{"web-dev", "js"}},
		{"JAVA", 253, "Offline", "abceg123", []string{"fullstack", "html"}},
		{"Py", 253, "Online", "abc8793", nil},
	}

	// paccking this data as JSON data
	// Marshal Indent is for beautify
	finalJson, err := json.MarshalIndent(Georgian, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodingJson() {
	jsonDataFromWeb := []byte(`
	
	
	{
		"Cousrename": "GO",
		"Price": 253,
		"Delivery-Mode": "Online",
		"tags": [
			"web-dev",
			"js"
		]
	}
	
	`)

	var storage course

	chechValid := json.Valid(jsonDataFromWeb)
	if chechValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDataFromWeb, &storage)
		fmt.Printf("%#v\n", storage)
	} else {
		fmt.Println("JSON WAS NOT VALID ")
	}

	// some times where just want to add data to key value
	// interface is if we dont know the type of data

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key: %v   value: %v    type: %T \n", key, value, value)
	}

}
