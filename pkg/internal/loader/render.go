package loader

import (
	"errors"
	"os"
	"path"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func Render(fn PlotFunc, opts Options) error {
	plt, err := plot.New()
	if err != nil {
		return err
	}
	err = fn(plt)
	if err != nil {
		return err
	}
	return writeChart(plt, opts)
}

func writeChart(plt *plot.Plot, opts Options) error {
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
