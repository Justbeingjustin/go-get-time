package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func main() {
	resp, err := http.Get("http://worldclockapi.com/api/json/Pst/now")
	if err != nil {
		println("There was an error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(string(body))
	theCurrentDateTime := gjson.Get(string(body), "currentDateTime")
	println("Current Date: ", theCurrentDateTime.String())

}
