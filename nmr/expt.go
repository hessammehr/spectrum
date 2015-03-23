package nmr

type Expt struct {
	Name   string
	Number int
	SW     float64
	O1P    float64
	NS     int
	TD     int
	FID    []int32
	FT     []complex128
	Phased []float64
	Shifts []float64
}
