package main

import (
	"fmt"

	"example.com/workhours/config"
)

func main() {
	fmt.Println("Hello World")
	cfg := config.LoadConfig()
	fmt.Println(cfg)
}
