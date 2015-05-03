package nmr

type Expt struct {
	// Book keeping
	Name   string
	Number int
	// Acquisition parameters
	SW  float64
	O1P float64
	NS  int
	TD  int
	// Processing parameters
	Ph0 float64
	Ph1 float64
	// Processed data
	FID    []float64
	FT     []complex128
	Phased []float64
	Shifts []float64
}
