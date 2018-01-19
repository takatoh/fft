package fft

import (
	"math"
	"math/cmplx"
)

func Fft(a []complex128, n int) []complex128 {
	x := make([]complex128, n)
	copy(x, a)

	j := 0
	for i := 0; i < n; i++ {
		if i < j {
			temp := x[j]
			x[j] = x[i]
			x[i] = temp
		}
		m := n / 2
		for {
			if j < m { break }
			j = j - m
			m = m / 2
			if m < 2 { break }
		}
		j = j + m
	}
	kmax := 1
	for {
		if kmax >= n {
			return x
		}
		istep := kmax * 2
		for k := 0; k < kmax; k++ {
			theta := complex(0.0, -1.0 * math.Pi * float64(k) / float64(kmax))
			for i := k; i < n; i += istep {
				j := i + kmax
				temp := x[j] * cmplx.Exp(theta)
				x[j] = x[i] - temp
				x[i] = x[i] + temp
			} 
		}
		kmax = istep
	}
}
