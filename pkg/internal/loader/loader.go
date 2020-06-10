package loader

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"github.com/fsnotify/fsnotify"
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
	if opts.Watch {
		// load first when watching for changes
		err := load(opts)
		if err != nil {
			return err
		}
		return watch(opts)
	}
	return load(opts)
}

func isWrite(evt fsnotify.Event) bool {
	return evt.Op&fsnotify.Write == fsnotify.Write
}

func isRemove(evt fsnotify.Event) bool {
	return evt.Op&fsnotify.Remove == fsnotify.Remove
}

func watch(opts Options) error {
	log.Printf("watching script %s for new changes\n", opts.ScriptPath)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()
	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				log.Println(event, ok)
				if !ok {
					errCh <- nil
					return
				}
				// NOTE: vim is weird
				// https://github.com/fsnotify/fsnotify/issues/94
				if isWrite(event) || isRemove(event) {
					err := load(opts)
					if err != nil {
						log.Printf("error loading script: %s", err)
					}
					if isRemove(event) {
						// TODO: add a better mechanism to
						// wait for the file to move back again.
						time.Sleep(1 * time.Second)
						err = watcher.Add(opts.ScriptPath)
						if err != nil {
							errCh <- err
							return
						}
					}
				}
			case err, ok := <-watcher.Errors:
				log.Println("caught error")
				if !ok {
					errCh <- nil
					return
				}
				errCh <- err
			case <-sigCh:
				log.Println("caught interrupt")
				errCh <- nil
			}
		}
	}()
	err = watcher.Add(opts.ScriptPath)
	if err != nil {
		return err
	}
	return <-errCh
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
