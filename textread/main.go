// Prints the output of a text file
// Emulates 'cat' on Linux

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	var filename string		
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	} else {
		log.Fatal("Please provide the filename as the first argument.")
	}
	
	tf, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	// Copies data from file to the system output. This is written as
	// 'os.Stdout' which actually is a file that implements the
	// Writer interface, as required by the method signiture of 'io.Copy'
	io.Copy(os.Stdout, tf)
}
