package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, errResp := http.Get(url)
	handleError(errResp)
	defer response.Body.Close()

	contentBytes, errRead := ioutil.ReadAll(response.Body)
	content := string(contentBytes)

	handleError(errRead)

	fmt.Println(content)

}
