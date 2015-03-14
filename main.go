package main

import (
	"fmt"
	"image"
	"math/rand"

	_ "code.google.com/p/draw2d/draw2d"
	"github.com/andlabs/ui"
	"github.com/hessammehr/spectrum/nmr"
)

type float float64

type spectrum struct {
	Data   nmr.NMR
	hzoom  [2]float
	vzoom  [2]float
	status ui.Label
}

var w ui.Window
var s spectrum

func (s *spectrum) Paint(rect image.Rectangle) *image.RGBA {
	//f, _ := os.Open("a.png")
	//rgba, _, _ := image.Decode(f)
	s.status.SetText(fmt.Sprintf("Hi %v", rand.Int()))
	return image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{200, 200}})
}
func (s *spectrum) Mouse(me ui.MouseEvent)  {}
func (s *spectrum) Key(ke ui.KeyEvent) bool { return true }

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
	status := ui.NewLabel("Status")
	s.status = status
	disp := ui.NewArea(200, 200, &s)
	stack := ui.NewVerticalStack(disp, status)
	stack.SetStretchy(0)
	w = ui.NewWindow("NMR", 800, 500, stack)
	w.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	w.Show()
}
func main() {
	s.Data = *nmr.Process("fid")

	go ui.Do(initGUI)
	err := ui.Go()
	if err != nil {
		panic(err)
	}
}
