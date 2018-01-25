package main

import (
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// fmt.Println(string(data))

	io.Copy(os.Stdout, file)
}
