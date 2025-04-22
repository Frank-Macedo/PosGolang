package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	println(string(body))

}

func mainy() {
	c := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"Name": "Franklin"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

func mainx() {
	c := http.Client{}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
