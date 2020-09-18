package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/wrnrlr/polyline"
	"image"
	"image/color"
	"log"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(700)))
		a := new(App)
		if err := a.loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type App struct{}

func (a *App) loop(w *app.Window) error {
	var ops op.Ops
	var l []f32.Point
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			pointer.Rect(image.Rectangle{Max: gtx.Constraints.Min}).Add(gtx.Ops)
			for _, ev := range gtx.Events(a) {
				ev, ok := ev.(pointer.Event)
				if !ok {
					continue
				}
				switch ev.Type {
				case pointer.Press:
					l = []f32.Point{ev.Position}
				case pointer.Drag:
					l = append(l, ev.Position)
				}
			}
			pointer.InputOp{Tag: a, Grab: false, Types: pointer.Press | pointer.Drag | pointer.Release}.Add(gtx.Ops)
			op.Push(gtx.Ops).Pop()
			polyline.Draw(l, unit.Dp(5).V, color.RGBA{255, 0, 0, 255}, gtx)
			e.Frame(gtx.Ops)
		}
	}
}
