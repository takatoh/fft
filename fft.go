package fft

import (
	"math"
	"math/cmplx"
)

func FFT(x []complex128, n int) []complex128 {
	y := fft(x, n)
	complex_n := complex(float64(n), 0.0)
	for i := 0; i < n; i++ {
		y[i] = y[i] / complex_n
	}
	return y
}

func IFFT(x []complex128, n int) []complex128 {
	y := make([]complex128, n)
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(x[i])
	}
	y = fft(y, n)
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(y[i])
	}
	return y
}

func FFTFreq(n int, dt float64) []float64 {
	ndt := float64(n) * dt
	f := make([]float64, n)
	for i := 0; i < n/2; i++ {
		f[i] = float64(i) / ndt
		f[n/2+i] = -float64(n/2-i) / ndt
	}
	return f
}

func RFFTFreq(n int, dt float64) []float64 {
	ndt := float64(n) * dt
	f := make([]float64, n/2+1)
	for i := 0; i < n/2+1; i++ {
		f[i] = float64(i) / ndt
	}
	return f
}

func PowerOf2(n int) int {
	nn := 2
	for nn < n {
		nn *= 2
	}
	return nn
}

func MakeComplexData(data []float64) ([]complex128, int) {
	n := len(data)
	nn := PowerOf2(n)
	cmplxData := make([]complex128, nn)
	for i := 0; i < n; i++ {
		cmplxData[i] = complex(data[i], 0.0)
	}
	for i := n; i < nn; i++ {
		cmplxData[i] = complex(0.0, 0.0)
	}
	return cmplxData, nn
}

func fft(a []complex128, n int) []complex128 {
	x := make([]complex128, n)
	copy(x, a)

	j := 0
	for i := 0; i < n; i++ {
		if i < j {
			x[i], x[j] = x[j], x[i]
		}
		m := n / 2
		for {
			if j < m {
				break
			}
			j = j - m
			m = m / 2
			if m < 2 {
				break
			}
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
			theta := complex(0.0, -1.0*math.Pi*float64(k)/float64(kmax))
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
