package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("hello world")
	//PerformGetRequest()
	//PostJsonRequest()
	PostFormRequest()
}

// store a host link in variable
// pass that link in http.Get(variable)  with error check
// always close at last better to give differ statenent
// cont, _ := ioutil.ReadAll(responce.Body)
// fmt.Println(string(cont))
// Another way is to by strings library by using [strings.Builder]

func PerformGetRequest() {
	const myhost = "http://localhost:3000"

	responce, err := http.Get(myhost)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	fmt.Println("Status code: ", responce.Status)
	fmt.Println("Length: ", responce.ContentLength)

	// cont, _ := ioutil.ReadAll(responce.Body)
	// fmt.Println(string(cont))

	// Another way is to by strings library

	var responceValue strings.Builder
	// its gives length of
	cont, _ := ioutil.ReadAll(responce.Body)
	byteCount, _ := responceValue.Write(cont)

	fmt.Println("Byte count: ", byteCount)
	fmt.Print(responceValue.String())

}

// store a host link in variable
// {
//create request body as requestBody := strings.NewReader(` {  post information as   "key":"value",.........     }  `)
// pass that link in http.Post(url link, "Application from where given in header", what you want to post [requestBody])  with error check
//  }
// always close at last better to give differ statenent
// cont, _ := ioutil.ReadAll(responce.Body)
// fmt.Println(string(cont))
// Another way is to by strings library by using [strings.Builder]

func PostJsonRequest() {

	const postUrl = "http://localhost:3000/postform"

	// fake json payload

	requestBody := strings.NewReader(`
		{   
			"coursename":"Golang",
			"price":"0",
			"platform":"www.google.com"

		}

`)

	responce, err := http.Post(postUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	cont, _ := ioutil.ReadAll(responce.Body)
	fmt.Println(string(cont))

	var responceValue strings.Builder
	databyte, _ := responceValue.Write(cont)
	fmt.Println(responceValue.String())
	fmt.Print("_", databyte)

}

func PostFormRequest() {

	const postformUrl = "http://localhost:3000/post"

	dataUrl := url.Values{}
	dataUrl.Add("fName", "Dhruvin")
	dataUrl.Add("lName", "Gadhiya")
	dataUrl.Add("email", "xyz@.gmail.co")

	responce, err := http.PostForm(postformUrl, dataUrl)
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	var valueofResponce strings.Builder

	content, _ := ioutil.ReadAll(responce.Body)

	dataByte, _ := valueofResponce.Write(content)

	fmt.Println("Byte count: ", dataByte)

	fmt.Println(valueofResponce.String())

}
