package poster

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
)

func SaveImageToFile(imageLink, fileName string) error {
    outFile, err := os.Create(fileName)
    if err != nil { return err }
    defer outFile.Close()

    image, err := getImage(imageLink)
    if err != nil { return err }

    err = jpeg.Encode(outFile, image, nil)
    return err
}

func getImage(url string) (image.Image, error) {
    client := http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil { return nil, nil }

    resp, err := client.Do(req)
    if err != nil { return nil, err }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK { return nil, fmt.Errorf("%s", resp.Status)}

    image, err := jpeg.Decode(resp.Body)
    if err != nil { return nil, err }

    return image, nil
}
