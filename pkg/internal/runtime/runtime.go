package runtime

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"gopkg.in/fsnotify.v1"
)

type handler interface {
	Start() error
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
	filepath   string
	filetype   string
	width      float32
	height     float32
	sigCh      chan os.Signal
	errCh      chan error
	evalInCh   chan evalOpts
	evalOutCh  chan evalCtx
}

func New(opts Options) *Runtime {
	rt := &Runtime{
		scriptPath: opts.ScriptPath,
		eval:       defaultEval,
		sigCh:      make(chan os.Signal),
		errCh:      make(chan error),
		evalInCh:   make(chan evalOpts),
		evalOutCh:  make(chan evalCtx),
	}
	if opts.Headless || opts.FilePath != "" {
		rt.handler = fileWriter{
			filePath: opts.FilePath,
			fileType: opts.FileType,
			height:   float32(opts.Height),
			width:    float32(opts.Width),
			evalCh:   rt.evalOutCh,
		}
	} else {
		rt.handler = &display{evalCh: rt.evalOutCh}
	}
	return rt
}

func (rt *Runtime) Run(ctx context.Context) error {
	signal.Notify(rt.sigCh, os.Interrupt)
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
	go Watch(rt.scriptPath, rt.evalInCh, rt.errCh)
	go rt.handler.Start()
	rt.evalOutCh <- c
	log.Println("waiting for file changes")
	for {
		select {
		case err := <-rt.errCh:
			return err
		case sig := <-rt.sigCh:
			fmt.Println(sig)
			return nil
		case opts := <-rt.evalInCh:
			now := time.Now()
			imgFn, err := rt.eval(ctx, opts)
			if err == nil {
				log.Println("script evaluated in", time.Since(now))
				rt.evalOutCh <- imgFn
			} else {
				log.Println(err)
				rt.evalOutCh <- evalCtx{err: err}
			}
		}
	}
}

func isWrite(evt fsnotify.Event) bool {
	return evt.Op&fsnotify.Write == fsnotify.Write
}

func isRemove(evt fsnotify.Event) bool {
	return evt.Op&fsnotify.Remove == fsnotify.Remove
}

func wait(path string, timeout time.Duration) error {
	start := time.Now()
	for {
		if time.Since(start) >= timeout {
			return fmt.Errorf("timeout")
		}
		_, err := os.Stat(path)
		if err == nil {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func clear() {
	// TODO: it would be nice if yaeji took
	// an io.Writer for stdout/stderr
	fmt.Fprintf(os.Stdout, "\033[H\033[2J")
	fmt.Fprintf(os.Stdout, "Program Output:\n\n")
}
