package main

import (
    "net/http"
    "sync"
    "image"
    "image/png"
    "go-study/mandelbrot/mdb"
)

func handler(w http.ResponseWriter, r *http.Request) {
    xmin, xmax, ymin, ymax := configure(r)
    width, height          := 2048, 2048

    img := image.NewRGBA(image.Rect(0, 0, width, height))

    var wg sync.WaitGroup
    for py := 0; py < height; py++ {
        y := float64(py) / float64(height) * (ymax-ymin) + ymin

        for px := 0; px < width; px++ {
            wg.Add(1)
            go func(px, py int) {
                x := float64(px) / float64(width) * (xmax-xmin) + xmin

                z := complex(x, y)

                img.Set(px, py, mdb.Mandelbrot(z))
                wg.Done()
            }(px, py)
        }
    }

    wg.Wait()
    png.Encode(w, img)
}
