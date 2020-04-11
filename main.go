package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"

	"github.com/otiai10/conway-app-go/conway"
	"github.com/otiai10/conway-app-go/presentation"
)

var (
	// World width
	width int
	// World height
	height int
	// Interval of generation alter
	interval int
)

func init() {
	flag.IntVar(&width, "w", 40, "Width of your world")
	flag.IntVar(&height, "h", 20, "Height of your world")
	flag.IntVar(&interval, "s", 200, "Interval of alternation (milli sec)")
}

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
	renderer := presentation.Renderer{Out: os.Stdout, Alive: "*", Dead: " "}
	world := conway.InitWorld(width, height)
	renderer.Update(world)
	for next := world.Next(); ; next = next.Next() {
		time.Sleep(time.Duration(interval) * time.Millisecond)
		renderer.Update(next)
	}
	return nil
}
