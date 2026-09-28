package main

import (
	"fmt"
	"os"

	"github.com/tekscape-inc/sample-legacy-go/internal/app/legacy"
)

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(legacy.Banner(name))
}
