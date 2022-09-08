package xkcd

var baseLinkLeft  string = "https://www.xkcd.com/"
var baseLinkRight string = "/info.0.json"

type Comic struct {
    Number      int     `json:"num"`
    Title       string  `json:"title"`
    Transcript  string  `json:"transcript"`
    URL         string  `json:"img"`
}

