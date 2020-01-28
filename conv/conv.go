package conv

// Float64to32 -
func Float64to32(d []float64) (o []float32) {
	o = make([]float32, len(d))
	for i := range d {
		o[i] = float32(d[i])
	}
	return
}

// Float64toInt16 -
func Float64toInt16(d []float64) (o []int16) {
	o = make([]int16, len(d))
	for i := range d {
		o[i] = int16(d[i])
	}
	return
}
