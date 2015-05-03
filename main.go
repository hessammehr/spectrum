package main

import (
	"errors"
	"image"
	"log"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/vg/vgimg"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"

	"github.com/hessammehr/spectrum/nmr"
)

var (
	width  = 800
	height = 500
	expt   nmr.Expt
)

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

func appMain(driver gxui.Driver) {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	p, _ := plot.New()
	xys, _ := XYs(expt.Shifts, expt.Phased)
	l, _ := plotter.NewLine(plotter.XYs(xys))
	p.Add(l)
	da := plot.MakeDrawArea(vgimg.NewImage(rgba))
	p.Draw(da)

	theme := dark.CreateTheme(driver)

	img := theme.CreateImage()
	img.SetTexture(driver.CreateTexture(rgba, 1))

	label := theme.CreateLabel()
	label.SetText("Status")

	layout := theme.CreateLinearLayout()
	layout.AddChild(img)
	layout.AddChild(label)
	layout.SetSizeMode(gxui.Fill)

	window := theme.CreateWindow(800, 600, "NMR")
	window.SetScale(1.0)
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	expt, err := nmr.ReadBruker("ma-catalyst/1")
	if err != nil {
		log.Panic("Failed to open experiment.")
	}
	expt.FFT()
	expt.Phase(0.1, 0.01)

	gl.StartDriver(appMain)
}
