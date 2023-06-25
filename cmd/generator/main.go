package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/terrylin13/gin-restful-generator/internal/generator"
)

func main() {
	project := flag.String("project", "", "name of the resource")
	flag.Parse()

	if *project == "" {
		fmt.Println("Missing name argument. Usage: go run main.go -project <project-name>")
		os.Exit(1)
	}

	generator.BaseDirectory(*project)
	fmt.Printf("Generated  %s  Success \n", *project)
}
