package poster

import (
	"fmt"
	"image/jpeg"
	"net/http"
    "os"
)

func SaveImageToFile(imageLink, fileName string) error {
    // Prepare file
    outFile, err := os.Create(fileName)
    if err != nil { return err }
    defer outFile.Close()

    // Prepare image
    client := http.Client{}
    req, err := http.NewRequest("GET", imageLink, nil)
    if err != nil { return err }

    resp, err := client.Do(req)
    if err != nil { return err }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK { return fmt.Errorf("%s", resp.Status)}

    image, err := jpeg.Decode(resp.Body)
    if err != nil { return err }

    // Write image to file
    err = jpeg.Encode(outFile, image, nil)
    return err
}
