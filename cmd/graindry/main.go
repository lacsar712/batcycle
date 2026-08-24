package main

import (
	"os"

	"github.com/lacsar712/batcycle/internal/app"
)

func main() {
	os.Exit(app.RunCLI())
}
