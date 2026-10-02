package main

import "os"

func main() {
	app := app.New()
	os.Exit(app.Run(os.Args))
}
