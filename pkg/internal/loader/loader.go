package loader

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"github.com/kevinschoon/qviz/pkg/internal/loader/symbols"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

type Options struct {
	ScriptPath string
	FileType   string
	FilePath   string
	// inches
	Height int
	Width  int
	Watch  bool
}

func DefaultOptions() *Options {
	return &Options{
		FilePath: "/dev/stdout",
		Height:   5,
		Width:    5,
	}
}

type PlotFunc func(*plot.Plot) error

func Load(opts Options) error {
	return load(opts)
}

func load(opts Options) error {
	fp, err := os.Open(opts.ScriptPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	fn, err := eval(fp)
	if err != nil {
		return err
	}
	return render(fn, opts)
}

func eval(reader io.Reader) (PlotFunc, error) {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(symbols.Symbols)
	raw, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	_, err = i.Eval(string(raw))
	if err != nil {
		return nil, err
	}
	qFn, err := i.Eval("QViz")
	if err != nil {
		return nil, err
	}
	pFn, ok := qFn.Interface().(func(*plot.Plot) error)
	if !ok {
		return nil, errors.New("bad qviz file")
	}
	return pFn, nil
}

func render(fn PlotFunc, opts Options) error {
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
