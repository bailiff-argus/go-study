package mdb

import (
    "image/color"
    "math"
    "math/cmplx"
)

func Mandelbrot(z complex128) color.Color {
    const iterations = 255
    const contrast = 4

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        mod := math.Pow(float64(n) / 255, 0.3) * 255
        if cmplx.Abs(v) > 2 {
            return color.RGBA{
                0   + uint8(mod),
                71  + uint8(mod * (255 - 71)  / 255),
                171 + uint8(mod * (255 - 171) / 255),
                255,
            }
        }
    }

    return color.Black
}
