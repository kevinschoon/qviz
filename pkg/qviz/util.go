package qviz

import (
	"io"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type VizFunc func(*plot.Plot) error

func SetTitle(title string) VizFunc {
	return func(plt *plot.Plot) error {
		plt.Title.Text = title
		return nil
	}
}

func SetLinePoints(label string, data [][2]float64) VizFunc {
	return func(plt *plot.Plot) error {
		xys := make(plotter.XYs, len(data))
		for i := 0; i < len(data); i++ {
			xys[i] = plotter.XY{
				X: data[i][0],
				Y: data[i][1],
			}
		}
		return plotutil.AddLinePoints(plt, label, xys)
	}
}

func SavePlot(w io.Writer) VizFunc {
	return func(plt *plot.Plot) error {
		out, err := plt.WriterTo(10*vg.Inch, 8*vg.Inch, "svg")
		if err != nil {
			return err
		}
		_, err = out.WriteTo(w)
		return err
	}
}

func SineDemo() VizFunc {
	fn := plotter.NewFunction(func(x float64) float64 { return float64(rand.Intn(10))*math.Sin(x) + 50 })
	return func(plt *plot.Plot) error {
		plt.Add(fn)
		plt.X.Min = 0
		plt.X.Max = 10
		plt.Y.Min = 0
		plt.Y.Max = 100
		return nil
	}
}
