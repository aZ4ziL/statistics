package statistics

// List type
type list []float64

// Single data is data that is presented simply and the data has not been
// arranged or grouped into interval classes. Grouped data is data that is
// usually presented in the form of a frequency table and the data has been
// compiled or grouped into interval classes.
type singleData struct {
	Data list // []float64
}

func NewSingleData(data list) *singleData {
	return &singleData{data}
}

// GetMinMax Finding the smallest and largest values ​​from a single data array
func (s singleData) GetMinMax() (min, max float64) {
	for i, e := range s.Data {
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
func (s singleData) Range() float64 {
	min, max := s.GetMinMax()
	return max - min
}

// Sum The total number of elements of an array
func (s singleData) Sum() float64 {
	var result float64
	for _, v := range s.Data {
		result += v
	}
	return result
}

// Average Finding the average value of a single data
func (s singleData) Average() float64 {
	n := len(s.Data)
	sum := s.Sum()

	return sum / float64(n)
}
