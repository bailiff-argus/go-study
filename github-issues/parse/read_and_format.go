package parse

import (
	// "fmt"
	"bufio"
	"html/template"
	"os"
	"strings"
	"time"

	"go-study/github-issues/github"
)


const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | setTimeTranche}}`


func DisplayResult (result *github.IssuesSearchResult) error {
    report, err := template.New("report").
        Funcs(template.FuncMap{"daysAgo": daysAgo}).
        Parse(templ)
    if err != nil { return err }

    err = report.Execute(os.Stdout, result)
    return err
}

func ReceiveInput (reader *bufio.Reader) (string, error) {
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }

    input = strings.ToLower(
        input[:len(input)-1],
    )

    return input, nil
}

func daysAgo(t time.Time) int {
    return int(time.Since(t).Hours() / 24)
}
