package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com/webhp?hl=en&sa=X&ved=0ahUKEwi_u8204dn9AhVTkYkEHeuwDTsQPAgI"

func main() {

	fmt.Println("URLS")
	fmt.Println(myurl)

	//parsing

	reault, _ := url.Parse(myurl)

	fmt.Println(reault.Scheme)
	fmt.Println(reault.RawQuery)
	fmt.Println(reault.Path)
	fmt.Println(reault.Host)
	fmt.Println(reault.Port())

	queryparameters := reault.Query()
	fmt.Printf("type is : %T\n", queryparameters)
	// Type of query parama=eters is key and value
	// so we can mainpulate as we want

	fmt.Println(queryparameters["hl"])

	partsofUrl := &url.URL{

		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)

}
