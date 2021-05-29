package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	//if the error is not nil (nil is returned when there are no errors)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// one way of printing out the html from the response

	// creating an empty byte slice with space fpr 99999 elements
	// bs := make([]byte, 99999)
	// Body has a Read function and we pass in the byte slice
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// Another way that we can print out the response
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wront this many bytpes:", len(bs))
	return len(bs), nil
}
