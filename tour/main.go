package main

// go run main.go word  -s=asdasd -m=1
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
