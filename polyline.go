package polyline

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image/color"
	"math"
)

const (
	rad90 = float32(90 * math.Pi / 180)
)

func Draw(points []f32.Point, width float32, col color.RGBA, gtx layout.Context) {
	length := len(points)
	for i, p := range points {
		drawCircle(p, width, col, gtx)
		if i < length-1 {
			drawLine(p, points[i+1], width, col, gtx)
		}
	}
}

func drawCircle(p f32.Point, radius float32, col color.RGBA, gtx layout.Context) {
	d := radius * 2
	const k = 0.551915024494 // 4*(sqrt(2)-1)/3
	defer op.Push(gtx.Ops).Pop()
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(f32.Point{X: p.X + radius, Y: p.Y})
	path.Cube(f32.Point{X: 0, Y: radius * k}, f32.Point{X: -radius + radius*k, Y: radius}, f32.Point{X: -radius, Y: radius})    // SE
	path.Cube(f32.Point{X: -radius * k, Y: 0}, f32.Point{X: -radius, Y: -radius + radius*k}, f32.Point{X: -radius, Y: -radius}) // SW
	path.Cube(f32.Point{X: 0, Y: -radius * k}, f32.Point{X: radius - radius*k, Y: -radius}, f32.Point{X: radius, Y: -radius})   // NW
	path.Cube(f32.Point{X: radius * k, Y: 0}, f32.Point{X: radius, Y: radius - radius*k}, f32.Point{X: radius, Y: radius})      // NE
	path.End().Add(gtx.Ops)
	box := f32.Rectangle{Min: f32.Point{X: p.X - radius, Y: p.Y - radius}, Max: f32.Point{X: p.X + d, Y: p.Y + d}}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: box}.Add(gtx.Ops)
}

func drawLine(p1, p2 f32.Point, width float32, col color.RGBA, gtx layout.Context) {
	tilt := angle(p1, p2)
	a := offsetPoint(p1, width, tilt+rad90)
	b := offsetPoint(p2, width, tilt+rad90)
	c := offsetPoint(p2, -width, tilt+rad90)
	d := offsetPoint(p1, -width, tilt+rad90)
	defer op.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(a)
	path.Line(b.Sub(a))
	path.Line(c.Sub(b))
	path.Line(d.Sub(c))
	path.Line(a.Sub(d))
	path.End().Add(gtx.Ops)
	box := boundingBox([]f32.Point{a, b, c, d})
	paint.PaintOp{Rect: box}.Add(gtx.Ops)
}

func boundingBox(points []f32.Point) (box f32.Rectangle) {
	for _, p := range points {
		box.Min.X = min(box.Min.X, p.X)
		box.Min.Y = min(box.Min.Y, p.Y)
		box.Max.X = max(box.Max.X, p.X)
		box.Max.Y = max(box.Max.Y, p.Y)
	}
	return box
}

func offsetPoint(point f32.Point, distance, angle float32) f32.Point {
	x := point.X + distance*cos(angle)
	y := point.Y + distance*sin(angle)
	return f32.Point{X: x, Y: y}
}

func angle(p1, p2 f32.Point) float32 {
	return atan2(p2.Y-p1.Y, p2.X-p1.X)
}

func cos(v float32) float32 {
	return float32(math.Cos(float64(v)))
}

func sin(v float32) float32 {
	return float32(math.Sin(float64(v)))
}

func min(x, y float32) float32 {
	return float32(math.Min(float64(x), float64(y)))
}

func max(x, y float32) float32 {
	return float32(math.Max(float64(x), float64(y)))
}

func atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}
