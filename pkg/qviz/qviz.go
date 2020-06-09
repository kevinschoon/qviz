package qviz

import (
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type Options struct{}

type Viz struct {
	Title  string
	Labels []string
	Data   [][][]float64
}

func Apply(fns ...VizFunc) error {
	rand.Seed(time.Now().Unix())

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}

	for _, fn := range fns {
		err = fn(plt)
		if err != nil {
			return err
		}
	}
	return nil

}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
