package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	// I'm not going to check the error here!
	req, err := http.NewRequest("GET", "http://example.com", nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// I'm not going to close the response body causing memory leak
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))}
