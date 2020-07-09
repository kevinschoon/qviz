package runtime

import (
	"bytes"
	"context"
	"image"
	"io"
	"log"
	"reflect"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"github.com/kevinschoon/qviz/pkg/internal/runtime/symbols"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type evalCtx struct {
	plot *plot.Plot
	err  error
}

type evalOpts struct {
	ScriptContents string
}

type evalFunc func(context.Context, evalOpts) (evalCtx, error)

func defaultEval(ctx context.Context, opts evalOpts) (evalCtx, error) {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(symbols.Symbols)
	var (
		// use a channel to halt program
		// execution after the first call to Render()
		done  = make(chan evalCtx)
		errCh = make(chan error)
	)
	// overlay closures here as what is defined in the
	// stdlib. any changes there must be reflected here
	// and vice versa.
	i.Use(map[string]map[string]reflect.Value{
		"github.com/kevinschoon/qviz/pkg/stdlib": {
			"New": reflect.ValueOf(func() *plot.Plot {
				plt, err := plot.New()
				if err != nil {
					panic(err)
				}
				return plt
			}),
			"Render": reflect.ValueOf(func(plt *plot.Plot) error {
				done <- evalCtx{plot: plt}
				return nil
			}),
			"Maybe": reflect.ValueOf(func(err error) {
				if err != nil {
					// TODO: maybe there is some way to
					// display the line this was called from?
					log.Println("QViz Encountered a Fatal Error:")
					errCh <- err
				}
			}),
		},
	})
	go func() {
		_, err := i.EvalWithContext(context.Background(), opts.ScriptContents)
		if err != nil {
			errCh <- err
		}
	}()
	select {
	case img := <-done:
		return img, nil
	case err := <-errCh:
		return evalCtx{}, err
	}
}

func toImage(ctx evalCtx, w, h float32) image.Image {
	img := vgimg.New(vg.Length(w), vg.Length(h))
	ctx.plot.Draw(draw.New(img))
	return img.Image()
}

func toReader(ctx evalCtx, w, h float32, format string) (io.Reader, error) {
	c, err := draw.NewFormattedCanvas(vg.Length(w), vg.Length(h), format)
	if err != nil {
		return nil, err
	}
	ctx.plot.Draw(draw.New(c))
	buf := bytes.NewBuffer(nil)
	c.WriteTo(buf)
	return buf, nil
}
