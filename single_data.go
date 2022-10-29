package statistics

// List type
type list []float64

// Single data is data that is presented simply and the data has not been
// arranged or grouped into interval classes. Grouped data is data that is
// usually presented in the form of a frequency table and the data has been
// compiled or grouped into interval classes.
type SingleData struct {
	Data list // []float64
}

// GetMinMax Finding the smallest and largest values ​​from a single data array
func (s SingleData) GetMinMax(data list) (min, max float64) {
	for i, e := range data {
		switch {
		case i == 0:
			max, min = e, e
		case e > max:
			max = e
		case e < min:
			min = e
		}
	}
	return min, max
}

// Range
// Calculating the difference between the smallest and largest values
func (s SingleData) Range(data list) float64 {
	min, max := s.GetMinMax(data)
	return max - min
}
