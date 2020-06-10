package loader

import (
	"io"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

type RenderOptions struct {
	FileType    string
	FilePath    string
	ChartWriter io.WriteCloser
	// inches
	Height int
	Width  int
}

func DefaultRenderOptions() *RenderOptions {
	return &RenderOptions{
		ChartWriter: os.Stdout,
		FileType:    "svg",
		Height:      5,
		Width:       5,
	}
}

func Render(ctx *PlotContext, opts RenderOptions) error {
	plt, err := plot.New()
	if err != nil {
		return err
	}
	err = ctx.PlotFunc(plt)
	if err != nil {
		return err
	}
	return writeChart(plt, opts)
}

func writeChart(plt *plot.Plot, opts RenderOptions) error {
	var fp io.WriteCloser
	if opts.FilePath != "" {
		w, err := os.Create(opts.FilePath)
		if err != nil {
			return err
		}
		fp = w
	} else {
		fp = opts.ChartWriter
	}
	defer fp.Close()
	w, err := plt.WriterTo(
		vg.Length(opts.Width)*vg.Inch, vg.Length(opts.Height)*vg.Inch, opts.FileType)
	if err != nil {
		return err
	}
	_, err = w.WriteTo(fp)
	return err
}
