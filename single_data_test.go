package statistics_test

import (
	"testing"

	"github.com/aZ4ziL/statistics"
)

var (
	data          []float64 = []float64{1, 2, 3, 4, 5}
	minShould     float64   = 1
	maxShould     float64   = 5
	rangeShould   float64   = 4
	sumShould     float64   = 15
	averageShould float64   = 3
)

func TestGetMinMax(t *testing.T) {
	s := statistics.NewSingleData(data)

	min, max := s.GetMinMax()
	if min != minShould {
		t.Errorf("The smallest value should be `%.2f` but the value that comes out is `%.2f`.", minShould, min)
		return
	}

	if max != maxShould {
		t.Errorf("The largest value should be `%.2f` but the value that comes out is `%.2f`.", maxShould, max)
		return
	}
}

func TestRange(t *testing.T) {
	s := statistics.NewSingleData(data)
	r := s.Range()

	if r != rangeShould {
		t.Errorf("The range value should be `%.2f` but the value that comes out it `%.2f`.", rangeShould, r)
		return
	}
}

func TestSum(t *testing.T) {
	s := statistics.NewSingleData(data)
	sum := s.Sum()

	if sum != sumShould {
		t.Errorf("The summary value should be `%.2f` but the value that comes out it `%.2f`.", sumShould, sum)
		return
	}
}

func TestAverage(t *testing.T) {
	s := statistics.NewSingleData(data)

	avg := s.Average()
	if avg != averageShould {
		t.Errorf("The average value should be `%.2f` but the value that comes out it `%.2f`.", averageShould, avg)
		return
	}
}
