package runtime

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
)

type handler interface {
	Start() error
	Update(evalCtx) error
}

type Options struct {
	FileType   string
	FilePath   string
	Height     int
	Width      int
	Headless   bool
	ScriptPath string
	Watch      bool
}

func DefaultOptions() *Options {
	return &Options{
		Width:  900,
		Height: 600,
	}
}

type Runtime struct {
	scriptPath string
	eval       evalFunc
	handler    handler
	headless   bool
	watch      bool
}

func New(opts Options) *Runtime {
	rt := &Runtime{
		scriptPath: opts.ScriptPath,
		eval:       defaultEval,
		watch:      opts.Watch,
	}
	if opts.Headless || opts.FilePath != "" {
		rt.handler = fileWriter{
			filePath: opts.FilePath,
			fileType: opts.FileType,
			height:   float32(opts.Height),
			width:    float32(opts.Width),
		}
	} else {
		rt.handler = &display{}
	}
	return rt
}

func (rt *Runtime) Run(ctx context.Context) error {
	var (
		sigCh  = make(chan os.Signal)
		evalCh = make(chan evalOpts)
		errCh  = make(chan error)
	)
	signal.Notify(sigCh, os.Interrupt)
	// read/eval the script once before
	// starting the loop
	var initErr error
	raw, err := ioutil.ReadFile(rt.scriptPath)
	if err != nil {
		initErr = err
		log.Println(err)
	}
	c, err := rt.eval(ctx, evalOpts{ScriptContents: string(raw)})
	if err != nil {
		initErr = err
		log.Println(err)
	}
	c.err = initErr
	if err := rt.handler.Start(); err != nil {
		return err
	}
	if rt.watch {
		go Watch(rt.scriptPath, evalCh, errCh)
		if err := rt.handler.Update(c); err != nil {
			return err
		}
		log.Println("waiting for file changes")
		for {
			select {
			case err := <-errCh:
				return err
			case sig := <-sigCh:
				fmt.Println(sig)
				return nil
			case opts := <-evalCh:
				now := time.Now()
				imgFn, err := rt.eval(ctx, opts)
				if err == nil {
					log.Println("script evaluated in", time.Since(now))
					if err := rt.handler.Update(imgFn); err != nil {
						return err
					}
				} else {
					log.Println(err)
					if err := rt.handler.Update(evalCtx{err: err}); err != nil {
						return err
					}
				}
			}
		}
	} else {
		return rt.handler.Update(c)
	}
}

func clear() {
	// TODO: it would be nice if yaeji took
	// an io.Writer for stdout/stderr
	fmt.Fprintf(os.Stdout, "\033[H\033[2J")
	fmt.Fprintf(os.Stdout, "Program Output:\n\n")
}
