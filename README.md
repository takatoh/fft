# fft

Fast Fourier Transform in Golang.

高速フーリエ変換/逆変換のライブラリです。

## Install

`go.mod` ファイルに次の行を追記してください。

```
require github.com/takatoh/fft v1.3.0
```

## Usage

フーリエ変換には `fft.FFT` 関数を使います。

```go
y := fft.FFT(x, n)
```

`x` がデータ（複素数）、`n` がデータ数。`n` は 2 のべき乗である必要があります。

逆変換には `fft.IFFT` 関数を使います。

```go
z := fft.IFFT(y, n)
```

フーリエ変換による各成分に対応する周波数を得るには `fft.FFTFreq` 関数または `fft.RFFTFreq` 関数を使います。`n` はデータ数（2のべき乗）、`dt` は時間間隔です。

```go
f := fft.FFTFreq(n, dt)
rf := fft.RFFTFreq(n, dt)
```

`fft.FFTFreq` は `n` 個、`fft.RFFTFreq` は `n/2+1` 個の周波数を返します。

`fft.FFTFreq` では、ナイキスト周波数 `f[n/2]` が負値です。この振る舞いは Python の `numpy` ライブラリにある `fft.fftfreq()` に倣いました。何故このようになっているのかはわかりません。

`fft.RFFTFreq` では、ナイキスト周波数 `rf[n/2]` が正値です。

## Utilities

`fft.FFT` や `fft.IFFT` に与えるデータは、2のべき乗個の複素数である必要があります。

`fft.PowerOf2` 関数は、引数 `n` よりも大きく、かつ最も小さい2のべき乗の数を返します。

```go
nn := fft.PowerOf2(n)
```

`fft.MakeComplexData` 関数は、実数のデータを複素数に変換し、かつ2のべき乗個に拡張します。

```go
complexData, nn := fft.MakeComplexData(realData)
```

## License

MIT License
