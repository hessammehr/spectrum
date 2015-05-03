package main

import (
	"errors"
	"fmt"
	"image"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/vg/vgimg"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"

	"github.com/hessammehr/spectrum/nmr"
)

type windowMode int

const (
	ph0Mode windowMode = iota
	ph1Mode
	hzoomMode
	vzoomMode
	intMode
	pickMode
	normalMode
)

var (
	width  = 800
	height = 500
	expt   *nmr.Expt
	redraw chan bool
	mode   windowMode
	shift  bool
)

type point struct{ X, Y float64 }

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

func scroll(event gxui.MouseEvent) {
	switch mode {
	case ph0Mode:
		amount := float64(event.ScrollY) / 20.0 / 360.0
		expt.Ph0 += amount
		expt.Phase(expt.Ph0, expt.Ph1)
		fmt.Println(expt.Ph0)
		redraw <- true
	case ph1Mode:
		amount := float64(event.ScrollY) / 50000.0 / 360.0
		expt.Ph1 += amount
		expt.Phase(expt.Ph0, expt.Ph1)
		fmt.Println(expt.Ph1)
		redraw <- true
	}
}

func keyDown(k gxui.KeyboardEvent) {
	fmt.Printf("Key down: %+v", k)
	switch k.Key {
	case 99, 103:
		shift = true
	case 34:
		if shift == false {
			mode = ph0Mode
		} else {
			mode = ph1Mode
		}

	case 1:
		mode = normalMode
	}
	fmt.Println(mode)
	redraw <- true
}

func keyUp(k gxui.KeyboardEvent) {
	switch k.Key {
	case 99, 103:
		shift = false
	default:
		mode = normalMode
	}
	fmt.Printf("Key up: %+v", k)
	redraw <- true
}

func appMain(driver gxui.Driver) {
	fmt.Println("Driver started")
	theme := dark.CreateTheme(driver)

	img := theme.CreateImage()
	img.OnMouseScroll(scroll)

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
	window.OnKeyDown(keyDown)
	window.OnKeyUp(keyUp)

	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	go func() {
		for {
			fmt.Println("Redrawing")

			p, _ := plot.New()
			xys, _ := XYs(expt.Shifts, expt.Phased)
			l, _ := plotter.NewLine(plotter.XYs(xys))
			p.Add(l)
			da := plot.MakeDrawArea(vgimg.NewImage(rgba))
			p.Draw(da)
			img.SetTexture(driver.CreateTexture(rgba, 1))

			fmt.Println("Done drawing!")
			fmt.Printf("%+v\n", mode)
			<-redraw
			// wait for redraw trigger
		}
	}()
	fmt.Println("I'm here!")
}

func main() {
	mode = normalMode
	redraw = make(chan (bool), 50)
	expt, _ = nmr.ReadBruker("ma-catalyst/1")
	// if err != nil {
	// 	log.Panic("Failed to open experiment.")
	// }
	expt.FFT()
	expt.Phase(expt.Ph0, expt.Ph1)
	gl.StartDriver(appMain)
}
