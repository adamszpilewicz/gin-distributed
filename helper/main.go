package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Foo struct {
	Bar []string
}

func main() {
	//foo1 := new(Foo) // or &Foo{}
	//getJson("http://example.com", foo1)
	//println(foo1.Bar)

	// alternately:

	foo2 := Foo{}
	getJson("https://raw.githubusercontent.com/adamszpilewicz/Building-Distributed-Applications-in-Gin/main/chapter02/recipes.json", &foo2)
	println(foo2.Bar)
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
