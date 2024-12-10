package main

import (
	"livecomments/cmd"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	
	if err := cmd.Execute(); err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
}
