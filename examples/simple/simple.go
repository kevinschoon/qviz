package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func QViz(plt *plot.Plot) error {
	plt.Title.Text = "An Example Plot"
	plotutil.AddLinePoints(plt, plotter.XYs{{X: 1, Y: 2}})
	return nil
}

func main() {
	p, _ := plot.New()
	QViz(p)
	p.Save(5*vg.Inch, 5*vg.Inch, "/tmp/out.png")
}
