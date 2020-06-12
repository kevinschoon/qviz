package main

import (
	"encoding/csv"
	"image/color"
	"io"
	"os"
	"strconv"

	qviz "github.com/kevinschoon/qviz/pkg/stdlib"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var invert = false
var windowSize = 5

var colors = map[string]color.Color{
	"primary":   color.Black,
	"secondary": color.White,
	"highlight": color.RGBA{R: 255, A: 255},
}

func sum(floats []float64) float64 {
	var n float64
	for _, float := range floats {
		n += float
	}
	return n
}

func main() {

	// assumes the script is being ran from the project root
	fp, err := os.Open("./examples/temperature/GlobalTemperatures.csv")
	qviz.Maybe(err)
	reader := csv.NewReader(fp)
	// skip the header
	reader.Read()
	i := 0
	temps := [][2]float64{}
	for {
		next, err := reader.Read()
		if err == io.EOF {
			break
		}
		// land average temperature
		lat, err := strconv.ParseFloat(next[7], 64)
		if err != nil {
			// skip missing data
			continue
		}
		temps = append(temps, [2]float64{float64(i), lat})
		i++
	}
	// convert values into XYs
	xys := make(plotter.XYs, len(temps))
	for i := 0; i < len(temps); i++ {
		xys[i].X = temps[i][0]
		xys[i].Y = temps[i][1]
	}
	scatter, _ := plotter.NewScatter(xys)
	scatter.Color = colors["primary"]
	plt := qviz.New()
	plt.Legend.Add("temperatures", scatter)
	// plot a sliding window
	xys = make(plotter.XYs, len(temps))
	var window []float64
	for i, temp := range temps {
		if len(window) < windowSize {
			window = append(window, temp[1])
			xys[i].X = float64(i)
			xys[i].Y = temp[1]
			continue
		}
		window = append(window[1:], temp[1])
		xys[i].X = float64(i)
		xys[i].Y = sum(window) / float64(windowSize)
	}
	sma, _ := plotter.NewLine(xys)
	sma.Color = colors["highlight"]
	sma.Dashes = []vg.Length{vg.Length(3), vg.Length(3)}
	plt.Legend.Add("SMA", sma)
	plt.Legend.Color = colors["primary"]
	grid := plotter.NewGrid()
	grid.Horizontal.Color = colors["primary"]
	grid.Vertical.Color = colors["primary"]
	plt.Title.TextStyle.Color = colors["primary"]
	plt.BackgroundColor = colors["secondary"]
	plt.Title.Color = colors["primary"]
	plt.Title.Text = "Average Daily Recorded Temperature Since 1750"
	plt.X.Min = 0
	plt.X.Max = float64(len(xys))
	plt.X.Label.Text = "Daily Temperature Recordings"
	plt.X.Label.Color = colors["primary"]
	plt.X.Tick.Color = colors["primary"]
	plt.X.LineStyle.Color = colors["primary"]
	plt.X.Color = colors["primary"]
	plt.X.Tick.Label.Color = colors["primary"]
	plt.Y.Min = 10
	plt.Y.Max = 20
	plt.Y.Label.Text = "Average Land Temperature (C)"
	plt.Y.Label.Color = colors["primary"]
	plt.Y.Tick.Color = colors["primary"]
	plt.Y.LineStyle.Color = colors["primary"]
	plt.Y.Color = colors["primary"]
	plt.Y.Tick.Label.Color = colors["primary"]
	plt.Add(scatter, sma, grid)
	qviz.Render(plt)
}

func init() {
	if invert {
		colors["primary"], colors["secondary"] =
			colors["secondary"], colors["primary"]
	}
}
