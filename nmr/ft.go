package nmr

import "github.com/jvlmdr/go-fftw/fftw"

func FloatToComplex(vals []float64) []complex128 {
	result := make([]complex128, len(vals))
	for i, val := range vals {
		result[i] = complex(val, 0)
	}
	return result
}

func FFT(sig []complex128) []complex128 {
	return fftw.FFT(&fftw.Array{Elems: sig[200:]}).Elems
}

func Reals(vals []complex128) []float64 {
	result := make([]float64, len(vals))
	for i, val := range vals {
		result[i] = real(val)
	}
	return result
}

func Imags(vals []complex128) []float64 {
	result := make([]float64, len(vals))
	for i, val := range vals {
		result[i] = imag(val)
	}
	return result
}
