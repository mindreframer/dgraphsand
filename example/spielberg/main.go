package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	baseURL = flag.String("d", "http://127.0.0.1:8080/query", "Dgraph server address")
)

var (
	allMichaelBay = `{
		director(func:allofterms(name, "michael bay")) {
			name@en
			director.film (orderdesc: initial_release_date) {
				name@en
				initial_release_date
			}
		}
	}`

	spielbergAfter1984 = `{
  director(func:allofterms(name, "steven spielberg")) {
    name@en
    director.film (orderasc: initial_release_date) @filter(ge(initial_release_date, "1960-08")) {
      name@en
      initial_release_date
    }
  }
}`

	spielberg90ties = `{
  director(func:allofterms(name, "steven spielberg")) {
    name@en
    director.film (orderasc: initial_release_date) @filter(ge(initial_release_date, "1990") AND le(initial_release_date, "2000")) {
      name@en
      initial_release_date
    }
  }
}`
	// @filter(anyofterms(name, "travis knight"))
	since2016 = `{
  films(func:ge(initial_release_date, "2016")) {
    name@en
    release: initial_release_date
    directed_by  {
      name@en
    }
  }
}`
)

func main() {
	res := execQuery(since2016)
	fmt.Println(string(res))
}

func execQuery(q string) []byte {
	client := &http.Client{}
	body := bytes.NewBufferString(q)
	req, err := http.NewRequest("POST", *baseURL, body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}
