package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/vg/vgimg"
	"github.com/andlabs/ui"
	"github.com/hessammehr/spectrum/nmr"
)

type float float64

var (
	width  = 800
	height = 500
)

type app struct {
	expt *nmr.Expt

	hzoom [2]float
	vzoom [2]float

	w ui.Window
	a ui.Area
	s ui.Label
}

var a app

func XYs(Xs []float64, Ys []float64) ([]struct{ X, Y float64 }, error) {
	if len(Xs) != len(Ys) {
		return nil, errors.New("Xs and Ys must have the same length.")
	}
	result := make([]struct{ X, Y float64 }, len(Xs))
	for i := range Xs {
		result[i] = struct{ X, Y float64 }{Xs[i], Ys[i]}
	}
	return result, nil
}

func (a *app) Paint(rect image.Rectangle) *image.RGBA {
	width, height = a.w.GetSize()
	height -= 50
	a.s.SetText(fmt.Sprintf("Hi %v, %v", width, height))
	a.a.SetSize(width, height)
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	p, _ := plot.New()
	xys, _ := XYs(a.expt.Shifts, a.expt.Phased)
	l, _ := plotter.NewLine(plotter.XYs(xys))
	p.Add(l)
	da := plot.MakeDrawArea(vgimg.NewImage(rgba))
	p.Draw(da)
	os.Remove("out.png")
	f, _ := os.Create("out.png")
	png.Encode(f, rgba)
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
	a.a = ui.NewArea(width, height, &a)
	stack := ui.NewVerticalStack(a.a, a.s)
	stack.SetStretchy(0)
	a.w = ui.NewWindow("NMR", width, height, stack)
	a.w.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	a.w.Show()
}
func main() {
	expt, _ := nmr.ReadBruker("ma-catalyst/1")
	expt.FFT()
	expt.Phase(0.1, 0.01)
	a.expt = expt
	go ui.Do(initGUI)
	err := ui.Go()
	if err != nil {
		panic(err)
	}
}
