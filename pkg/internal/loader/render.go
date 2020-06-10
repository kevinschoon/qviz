package loader

import (
	"errors"
	"os"
	"path"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

type RenderOptions struct {
	FileType string
	FilePath string
	// inches
	Height int
	Width  int
}

func DefaultRenderOptions() *RenderOptions {
	return &RenderOptions{
		FileType: "",
		FilePath: "/dev/stdout",
		Height:   5,
		Width:    5,
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
	var (
		ft string = opts.FileType
	)
	fp, err := os.Create(opts.FilePath)
	if err != nil {
		return err
	}
	if ft == "" {
		// use the extension to guess the file type
		ft = strings.Replace(path.Ext(opts.FilePath), ".", "", 1)
		if ft == "" {
			return errors.New("could not guess file type")
		}
	}
	defer fp.Close()
	w, err := plt.WriterTo(
		vg.Length(opts.Width)*vg.Inch, vg.Length(opts.Height)*vg.Inch, ft)
	if err != nil {
		return err
	}
	_, err = w.WriteTo(fp)
	return err
}
