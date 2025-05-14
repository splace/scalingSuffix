package scalingSuffix

import "math"

// suffix stores the number of scaling factors. (-ve scale down)
type Suffix int8

// calculate the suffix from the value and the scaling factor base
func SuffixFor(v,sf float64) Suffix{
	return Suffix(math.Floor(math.Log(math.Abs(v))*sf))
}

// calculate the suffices scaling from the scaling factor base
func (s Suffix) Scale(sf float64) float64{
	return math.Pow(sf,-float64(s)) // -ve exponent to Pow removes a divide.
}



