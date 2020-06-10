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

type PlotContext struct {
	PlotFunc   PlotFunc
	ScriptText string
}

type PlotFunc func(*plot.Plot) error

func Load(reader io.Reader) (*PlotContext, error) {
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
	return &PlotContext{
		PlotFunc:   pFn,
		ScriptText: string(raw),
	}, nil
}

func LoadPath(path string) (*PlotContext, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	return Load(fp)
}
