package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Getenv("apr"), os.Getenv("mart"))
	fmt.Print("Hello world 123")
	err := fmt.Errorf("%s/n", "Ашипка")
	log.Fatal(err)
}
