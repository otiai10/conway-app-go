package main

import (
	"io"
	"log"
	"os"
)

func main() {
	app := new(App)
	if err := app.Run(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}

// App ...
type App struct{}

// Run ...
func (app *App) Run(stdout io.Writer) error {
	return nil
}
