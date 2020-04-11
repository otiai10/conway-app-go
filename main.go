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
	// Expression of alive cell
	aliveEx string
	// Expression of dead cell
	deadEx string
)

func init() {
	flag.IntVar(&width, "w", 40, "Width of your world")
	flag.IntVar(&height, "h", 20, "Height of your world")
	flag.IntVar(&interval, "s", 200, "Interval of alternation (milli sec)")
	flag.StringVar(&aliveEx, "a", "*", "Expression string for alive cells")
	flag.StringVar(&deadEx, "d", " ", "Expression string for dead cells")
	flag.Parse()
}

// App ...
type App struct {
	WorldWidth          int
	WorldHeight         int
	AlternationInterval time.Duration
	ExpressionAlive     string
	ExpressionDead      string
}

// Run ...
func (app *App) Run(stdout io.Writer) error {
	renderer := presentation.Renderer{
		Out:   os.Stdout,
		Alive: app.ExpressionAlive,
		Dead:  app.ExpressionDead,
	}
	world := conway.InitWorld(width, height)
	renderer.Update(world)
	for next := world.Next(); ; next = next.Next() {
		time.Sleep(time.Duration(interval) * time.Millisecond)
		renderer.Update(next)
	}
	return nil
}

func main() {
	app := &App{
		WorldWidth:          width,
		WorldHeight:         height,
		AlternationInterval: time.Duration(interval) * time.Millisecond,
		ExpressionAlive:     aliveEx,
		ExpressionDead:      deadEx,
	}
	if err := app.Run(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
