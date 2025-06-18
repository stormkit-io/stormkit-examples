package main

import "log"

func main() {
	if _, err := Connection(); err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	StartServer()
}
