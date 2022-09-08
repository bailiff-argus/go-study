package parse

import (
    "fmt"

    "os"
    "os/exec"

    "io"
    "io/ioutil"

    "log"

    "encoding/json"

    "bytes"
)

var tempFile string = ".temp"

func GetInputFromEditor (initialContents string) string {
    editor := os.Getenv("EDITOR")
    file, err := os.Create(tempFile)

    if err != nil {
        log.Fatal(err)
    }

    fileName := file.Name()

    file.WriteString(initialContents)
    file.Close()

    openTemp := exec.Command(editor, fileName)
    openTemp.Stdin = os.Stdin
    openTemp.Stdout = os.Stdout
    openTemp.Stderr = os.Stderr
    openTemp.Run()

    contents, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    os.Remove(fileName)

    return string(contents)
}

func FormRequestBody (title string, text string) (io.Reader, error) {
    if text == "" {
        return nil, fmt.Errorf("empty issue description")
    }

    json, _ := json.Marshal(
        map[string]string{
            "title": title,
            "body":  text,
        },
    )
    
    bodyBuffer := bytes.NewBuffer(json)

    return bodyBuffer, nil
}
