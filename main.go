package main

import (
	"fmt"
	"image"
	"math/rand"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/vg/vgimg"
	"github.com/andlabs/ui"
	"github.com/hessammehr/spectrum/nmr"
)

type float float64

type app struct {
	expt *nmr.Expt

	hzoom [2]float
	vzoom [2]float

	w ui.Window
	a ui.Area
	s ui.Label
}

var a app

func (a *app) Paint(rect image.Rectangle) *image.RGBA {
	//f, _ := os.Open("a.png")
	//rgba, _, _ := image.Decode(f)
	a.s.SetText(fmt.Sprintf("Hi %v", rand.Int()))
	rgba := image.NewRGBA(image.Rect(0, 0, 800, 500))
	p, _ := plot.New()
	l, _ := plotter.NewLine(plotter.XYs{{0, 0}, {1, 1}, {2, 2}})
	p.Add(l)
	da := plot.MakeDrawArea(vgimg.NewImage(rgba))
	p.Draw(da)
	// gc := draw2d.NewGraphicContext(rgba)
	// gc.SetStrokeColor(image.Black)
	// gc.SetFillColor(image.White)
	// gc.Clear()
	// for i := 0.0; i < 360; i = i + 10 { // Go from 0 to 360 degrees in 10 degree steps
	// 	gc.BeginPath() // Start a new path
	// 	gc.Save()      // Keep rotations temporary
	// 	gc.MoveTo(144, 144)
	// 	gc.Rotate(i * (math.Pi / 180.0)) // Rotate by degrees on stack from 'for'
	// 	gc.RLineTo(72, 0)
	// 	gc.Stroke()
	// 	gc.Restore() // Get back the unrotated state
	// }
	// gc.Save()
	return rgba
}
func (a *app) Mouse(me ui.MouseEvent)  {}
func (a *app) Key(ke ui.KeyEvent) bool { return true }

func initGUI() {
	// b := ui.NewButton("Button")
	// c := ui.NewCheckbox("Checkbox")
	// tf := ui.NewTextField()
	// tf.SetText("Text Field")
	// pf := ui.NewPasswordField()
	// pf.SetText("Password Field")
	// l := ui.NewLabel("Label")
	// t := ui.NewTab()
	// t.Append("Tab 1", ui.Space())
	// t.Append("Tab 2", ui.Space())
	// t.Append("Tab 3", ui.Space())
	// g := ui.NewGroup("Group", ui.Space())
	a.s = ui.NewLabel("Status")
	a.a = ui.NewArea(200, 200, &a)
	stack := ui.NewVerticalStack(a.a, a.s)
	stack.SetStretchy(0)
	a.w = ui.NewWindow("NMR", 800, 500, stack)
	a.w.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	a.w.Show()
}
func main() {
	expt, _ := nmr.ReadBruker("ma-catalyst/1")
	a.expt = expt
	go ui.Do(initGUI)
	err := ui.Go()
	if err != nil {
		panic(err)
	}
}
