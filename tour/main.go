package main

import (
	"log"
	"tools/tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err :%v", err)
	}
}
