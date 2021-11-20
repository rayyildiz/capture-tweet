package main

import (
	"log"
	"os"
)

func main() {
	projectName := os.Getenv("GOOGLE_PROJECT_NAME")
	if len(projectName) < 1 {
		log.Fatalf("GOOGLE_PROJECT_NAME ıs required")
	}

}
