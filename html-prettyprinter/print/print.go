package print

import (
    "fmt"

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
            nodeString += "/>\n"
        } else {
            nodeString += ">\n"
            depth++
        }

        fmt.Println(nodeString)
    } else if n.Type == html.CommentNode {
        fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
    } else if n.Type == html.DoctypeNode {
        fmt.Printf("%*s<!%s>\n", depth*2, "", n.Data)
    } else if n.Type == html.TextNode {
        fmt.Println(n.Data)
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
