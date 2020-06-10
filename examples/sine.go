package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func QViz(plt *plot.Plot) error {
	plt.Add(plotter.NewFunction(func(x float64) float64 {
		return 10*math.Sin(x) + 50
	}))
	plt.X.Min = 0
	plt.X.Max = 10
	plt.Y.Min = 0
	plt.Y.Max = 100
	return nil
}
