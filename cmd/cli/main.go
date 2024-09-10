package main

import (
	"flag"
	"log"
	"os"
	"strings"

	create_date "allo/internal/allocator/create-date"
	raw_jpeg "allo/internal/allocator/raw-jpeg"
	"allo/internal/processor"
)

func main() {
	// log pwd of the current cli process
	osDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}
	log.Println("current working directory: ", osDir)

	stepsString := flag.String("steps", "", "comma separated list of allocating steps: \n- create-date \n- raw-jpeg")
	pathDir := flag.String("path", ".", "path to directory")

	flag.Parse()

	if *stepsString == "" {
		printWelcomeText()
		return
	}

	processor := processor.New()

	steps := strings.Split(*stepsString, ",")
	for _, step := range steps {
		switch strings.Trim(step, " ") {
		case "create-date":
			processor.AddAllocator(create_date.New(create_date.WithMode(create_date.ModeYear)))
			processor.AddAllocator(create_date.New(create_date.WithMode(create_date.ModeMonth)))
			processor.AddAllocator(create_date.New(create_date.WithMode(create_date.ModeDay)))
		case "raw-jpeg":
			processor.AddAllocator(raw_jpeg.New())
		default:
			log.Fatalf("unknown allocator: %s", step)
		}
	}

	processor.Run(*pathDir)

}

func printWelcomeText() {
	log.Println("\033[36m" + `
====================================
 _______  ___      ___      _______ 
|   _   ||   |    |   |    |       |
|  |_|  ||   |    |   |    |   _   |
|       ||   |    |   |    |  | |  |
|       ||   |___ |   |___ |  |_|  |
|   _   ||       ||       ||       |
|__| |__||_______||_______||_______|
---- an image allocator service ----

====================================
` + "\033[0m" + `
type "allo -h" for help`)
}
