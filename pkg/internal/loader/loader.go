package loader

import (
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"github.com/kevinschoon/qviz/pkg/internal/loader/symbols"
	"gonum.org/v1/plot"
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
	fp, err := os.Open(opts.ScriptPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	fn, err := eval(fp)
	if err != nil {
		return err
	}
	return Render(fn, opts)
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
