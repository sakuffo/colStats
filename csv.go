package main

func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

type statFunc func(data []float64) float64

// func csv2float(r io.Reader, column int) ([]float64, error) {

// }
