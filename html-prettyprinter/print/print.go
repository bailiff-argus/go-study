package print

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func StartElement(n *html.Node) {
    var nodeString string

    if n.Type == html.ElementNode {
        nodeString = fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
        for _, attr := range n.Attr {
            nodeString += fmt.Sprintf(" %s='%s'", attr.Key, attr.Val)
        }

        if n.FirstChild == nil {
            // nodeString += "/>\n"
            nodeString += "/>"
        } else {
            // nodeString += ">\n"
            nodeString += ">"
            depth++
        }

        fmt.Println(nodeString)
    } else if n.Type == html.CommentNode {
        fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
    } else if n.Type == html.DoctypeNode {
        fmt.Printf("%*s<!DOCTYPE %s>\n", depth*2, "", n.Data)
    } else if n.Type == html.TextNode {
        text := strings.TrimSpace(n.Data)
        for _, line := range strings.Split(text, "\n") {
            if line == "" {
                continue
            }

            fmt.Printf("%*s%s\n", depth*2, "", line)
        }
    }

}

func EndElement(n *html.Node) {
    if n.Type == html.ElementNode {
        if n.FirstChild == nil {
            return
        }

        depth--
        fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
    }
}
