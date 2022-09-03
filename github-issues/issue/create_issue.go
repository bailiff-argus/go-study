package issue

import (
    "fmt"
    "log"
    // "net/url"
    "net/http"
    "encoding/json"

    "bytes"
    "os"
    "os/exec"
    "io/ioutil"
)

var baseURL string = "https://api.github.com/repos/"
var tempFile string = ".temp"

func CreateIssue (title string, repo string, auth string) error {
    if auth == "" {
        return fmt.Errorf("cannot create issue without authorization")
    }

    // var body []byte
    // titleRQ := url.QueryEscape(title)
    // textRQ  := url.QueryEscape(getUserInput())

    titleRQ := title
    textRQ  := getUserInput()

    if textRQ == "" {
        return fmt.Errorf("empty issue description")
    }

    json, _ := json.Marshal(
        map[string]string{
            "title": titleRQ,
            "body":  textRQ,
        },
    )

    client := &http.Client{}
    link := formQueryLink(repo)
    bodyBuffer := bytes.NewBuffer(json)

    fmt.Println(bodyBuffer.Bytes())

    req, _ := http.NewRequest("POST", link, bodyBuffer)

    req.Header.Add("Accept", "application/vnd.github+json")

    authKey := "Authorization"
    authValue := fmt.Sprintf("Bearer %s", auth)
    req.Header.Set(authKey, authValue)

    fmt.Println(req.Header)
    fmt.Println(req.Body)

    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    if resp.StatusCode != http.StatusCreated {
        resp.Body.Close()
        return fmt.Errorf("issue creation failed: %s", resp.Status)
    }

    resp.Body.Close()

    return nil
}

func formQueryLink (repo string) string {
    return fmt.Sprintf("%s%s/issues", baseURL, repo)
}

func getUserInput () string {
    editor := os.Getenv("EDITOR")
    file, err := os.Create(tempFile)
    if err != nil {
        log.Fatal(err)
    }

    fileName := file.Name()
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
