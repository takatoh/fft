# fft

Fast Fourier Transform in Golang.

高速フーリエ変換/逆変換のライブラリです。

## Install
``` go get github.com/takatoh/fft```

## Usage
フーリエ変換には fft.FFT 関数を使います。
``` y := fft.FFT(x, n)```
x がデータ（複素数）、n がデータ数。n は 2 のべき乗である必要があります。

逆変換には fft.IFFT 関数を使います。
``` z := fft.IFFT(y, n)```

## License
MIT License
