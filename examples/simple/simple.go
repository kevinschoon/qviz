package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/brewer"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func QViz(plt *plot.Plot) error {
	plt.Title.Text = "QViz Example Plot"
	plt.Title.TextStyle.Font.Size = 5 * vg.Millimeter
	plt.X.Label.Text = "The X Axis..."
	plt.Y.Label.Text = "The Y Axis..."
	plt.X.Padding = vg.Points(5)
	palette, _ := brewer.GetPalette(brewer.TypeAny, "Dark2", 3)
	quad := plotter.NewFunction(func(x float64) float64 {
		return x * x
	})
	quad.Color = palette.Colors()[0]
	quad.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	quad.Width = vg.Points(2)
	plt.Legend.Add("x^2", quad)
	exp := plotter.NewFunction(func(x float64) float64 {
		return math.Pow(2, x)
	})
	exp.Color = palette.Colors()[1]
	exp.Width = vg.Points(2)
	plt.Legend.Add("2^x", exp)
	sin := plotter.NewFunction(func(x float64) float64 {
		return 10*math.Sin(x) + 100
	})
	sin.Color = palette.Colors()[2]
	sin.Width = vg.Points(2)
	plt.Legend.Add("10*sin(x)+100", sin)
	grid := plotter.NewGrid()
	plt.Add(grid, quad, exp, sin)
	plt.X.Min = 0
	plt.X.Max = 20
	plt.Y.Min = 0
	plt.Y.Max = 200
	return nil
}
