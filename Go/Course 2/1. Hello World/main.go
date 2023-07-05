package main

import (
	"fmt"
	"os"
)

const worldName = "Go"

func main() {
	shoutWorld(worldName)
}

func shoutWorld(worldName string) {
	fmt.Fprintf(os.Stdout, "Hello %s World!", worldName)
}
