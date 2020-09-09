package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// make function takes in the slice and the size, so that slice can grow or shrink up to the size specifed.
	// The reason for passing a lage size is so that when response is being written into the slice of byte, it
	// should be enough to store the body. Similarly at the time of reading from it.
	// Otherwise, the response will write only some portion of the body. Though it is not en efficient way of handing.

	// Method #1 for printing HTTP response body
	/*
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	// Method #2

	lg := logWriter{}

	io.Copy(lg, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("The number of bytes read and written is: ", len(bs))
	return len(bs), nil
}
