package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/markbates/pkger"
)

func main() {
	f, err := pkger.Open("/hello-world.txt")
	if err != nil {
		log.Fatal("could not open file", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatal("could not get stat of file", err)
	}

	fmt.Println("Name: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("ModTime: ", info.ModTime())

	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal("could not copy file to stdout", err)
	}
}
