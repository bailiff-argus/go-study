package main

import (
    "strconv"
    "log"
    "net/http"
    "net/url"
)

func pullFloatFromForm(form url.Values, name string, def float64) (float64) {
    var value float64
    var err error

    if val, ok := form[name]; ok {
        value, err = strconv.ParseFloat(val[0], 10)
        if err != nil {
            value = float64(def)
            log.Println(err)
        }
    } else {
        value = float64(def)
    }

    return value
}

func configure(r *http.Request) (float64, float64, float64, float64) {
    if err := r.ParseForm(); err != nil { log.Print(err) }
    form := r.Form

    var xCenter, yCenter, width float64

    xCenter = pullFloatFromForm(form, "xCenter", 0.)
    yCenter = pullFloatFromForm(form, "yCenter", 0.)
    width   = pullFloatFromForm(form, "width",   2.)

    xmin, xmax := xCenter - width, xCenter + width
    ymin, ymax := yCenter - width, yCenter + width

    return xmin, xmax, ymin, ymax

}
