package runtime

import (
	"image"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

type display struct {
	ctx    evalCtx
	evalCh <-chan evalCtx
}

func (d *display) Start() error {
	errCh := make(chan error)
	go func() {
		w := app.NewWindow(app.Title("QViz"))
		if err := d.loop(w); err != nil {
			errCh <- err
			return
		}
	}()
	app.Main()
	return <-errCh
}

func (d *display) Handle(ctx *evalCtx) error {
	return nil
}

func (d *display) handle(ops *op.Ops, evt event.Event) error {
	switch evt := evt.(type) {
	case system.DestroyEvent:
		return evt.Err
	case system.FrameEvent:
		gtx := layout.NewContext(ops, evt)
		var img image.Image
		x := float32(gtx.Constraints.Max.X) - 100
		y := float32(gtx.Constraints.Max.Y) - 100
		if d.ctx.err == nil {
			img = toImage(d.ctx, x, y)
		} else {
			img = loading(uint(x), uint(y))
		}
		// image display
		paint.NewImageOp(img).Add(gtx.Ops)
		paint.PaintOp{Rect: f32.Rect(0, 0, x, y)}.Add(gtx.Ops)
		evt.Frame(gtx.Ops)
	}
	return nil
}

func (d *display) loop(w *app.Window) error {
	material.NewTheme(gofont.Collection())
	ops := new(op.Ops)
	evtCh := w.Events()
	for {
		select {
		case evt := <-evtCh:
			err := d.handle(ops, evt)
			if err != nil {
				panic(err)
			}
		case evalCtx := <-d.evalCh:
			d.ctx = evalCtx
			w.Invalidate()
		}
	}
}
