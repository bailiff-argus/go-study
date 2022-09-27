package mdb

import (
    "math/cmplx"
    "image/color"
)

func Newton(z complex128) color.Color {
    const iterations = 200
    const contrast = 16

    for n := uint8(0); n < iterations; n++ {
        f_val := f(z)
        if cmplx.Abs(f_val) < 0.01 {
            return color.Gray{255 - n*contrast}
        }

        f_der := fder(z)
        z = z - f_val / f_der
    }
    return color.Black
}

func f(z complex128) complex128 {
    return cmplx.Pow(z, 4) - 1
}

func fder(z complex128) complex128 {
    return 4 * cmplx.Pow(z, 3) - 1
}
