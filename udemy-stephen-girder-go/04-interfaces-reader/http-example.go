package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func HttpExample() {
	// https://jsonplaceholder.typicode.com/posts

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println("Response", resp)
	// bareReadWriteImplementation(resp)
	// IoCopyExample(resp)
	CustomWriterExample(resp)

	// io.Copy(os.Stdout, resp.Body) // this is a more efficient way to do the same thing÷÷
}

func bareReadWriteImplementation(resp *http.Response) {
	// read from response body
	// write to os.Stdout
	// this is a very basic implementation of io.Reader and io.Writer

	// later we can use io.Copy to copy from one to another

	// Read from resp.Body takes in a byte slice on which it writes the data
	// it returns the number of bytes read and an error if any
	buf := make([]byte, 55000)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading response body:", err)
			}
			fmt.Printf("the value of n in the last read is: %d\n", n)
			break
		}
	}
	// fmt.Println(string(buf[:n]) + "...")
}

func IoCopyExample(resp *http.Response) {
	// io.Copy is a more efficient way to copy from one to another
	// Output of the response is handled in the writer function logic of os.Stdout
	// Input of the response is handled in the reader function logic of resp.Body
	res, err := io.Copy(os.Stdout, resp.Body)

	if err != nil {
		fmt.Println("Error copying response body:", err)
	} else {
		fmt.Printf("\nNumber of bytes copied: %d\n", res)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	// custom writer to log the number of bytes written
	fmt.Println("This is a custom log writer")
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %d\n", len(bs))
	return len(bs), nil
}

func CustomWriterExample(resp *http.Response) {
	lw := logWriter{}
	n, err := io.Copy(lw, resp.Body)
	if err != nil {
		fmt.Println("Error copying response body to custom writer:", err)
	} else {
		fmt.Printf("Number of bytes copied to custom writer: %d\n", n)
	}
}
