package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Joke struct {
	Joke string
}

func fetchJoke() Joke {
	response, err := http.Get("https://v2.jokeapi.dev/joke/Any?type=single")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var joke Joke;
	json.Unmarshal([]byte(responseData), &joke);
	return joke
}

func main() {
	var joke = fetchJoke()
	fmt.Println(joke.Joke)
}
