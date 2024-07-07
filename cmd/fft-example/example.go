package main

import (
	"fmt"

	"github.com/takatoh/fft"
)

func main() {
	x0 := []float64{
		5,
		32,
		38,
		-33,
		-19,
		-10,
		1,
		-8,
		-20,
		10,
		-1,
		4,
		11,
		-1,
		-7,
		-2,
	}
	x, nn := fft.MakeComplexData(x0)

	y := fft.FFT(x, nn)
	z := fft.IFFT(y, nn)
	f := fft.FFTFreq(nn, 0.02)

	fmt.Println(" K   DATA  FOURIER TRANSFORM  INVERSE TRANSFORM  FREQUENCY")
	for k := 0; k < nn; k++ {
		fmt.Printf("%2d %6.1f  %8.3f%8.3f   %8.3f%8.3f   %8.3f\n",
			k, x0[k], real(y[k]), imag(y[k]), real(z[k]), imag(z[k]), f[k])
	}
}
