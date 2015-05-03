package nmr

import "math"

func (e *Expt) Window() {

}

func (e *Expt) FFT() {
	e.FT = FFT(FloatToComplex(e.FID))
	e.Shifts = shifts(e.SW, e.O1P, e.TD)
}

func (e *Expt) Phase(ph0, ph1 float64) {
	for i := range e.FT {
		phase := ph0 + float64(i)*ph1
		e.FT[i] *= complex(math.Cos(phase), math.Sin(phase))
	}
	e.Phased = Reals(e.FT)[:e.TD/2]
}

func shifts(sw float64, o1p float64, td int) []float64 {
	points := td / 2
	result := make([]float64, points)
	min := o1p - sw/2
	step := sw / float64(points)
	for i := range result {
		result[i] = min + step*float64(i)
	}
	return result
}
