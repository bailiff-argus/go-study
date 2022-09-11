package omdb

const baseLink string = "https://www.omdbapi.com/"

type poster struct{
    ImgLink     string      `json:"Poster"`
}   
