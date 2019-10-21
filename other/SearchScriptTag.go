// This code make to search Script tag on static web page
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("write url")

	fmt.Println(" http transport error is :", err)
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("read error is ", err)

	fmt.Println(string(body))
}
