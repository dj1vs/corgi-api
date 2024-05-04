package html_basics

import (
	"golang.org/x/net/html"
)

func GetAttributes(n *html.Node, key string) ([]string, bool) {
	isFound := false
	attrs := []string{}

	for _, attr := range n.Attr {
		if attr.Key == key {
			isFound = true
			attrs = append(attrs, attr.Val)
		}
	}

	return attrs, isFound
}

func GetAttribute(n *html.Node, key string) (string, bool) {

	for _, attr := range n.Attr {

		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}

// children - how many levels deep should we go
func CollectText(n *html.Node, children int) string {
	text := ""
	if n.Type == html.TextNode {
		text += n.Data
		if n.Parent != nil && n.Parent.Data == "li" {
			text += "\n"
		}
	} else if n.Type == html.ElementNode && n.Data == "br" {
		text += "\n"
	}

	if children > 0 {
		children -= 1
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			text += CollectText(c, children)
		}
	}

	return text
}

func checkAttr(n *html.Node, attrName string, attrVal string) bool {

	if n.Type == html.ElementNode {

		s, ok := GetAttribute(n, attrName)
		if ok && s == attrVal {
			return true
		}
	}

	return false
}

func traverse(n *html.Node, attrName string, attrVal string) *html.Node {

	if checkAttr(n, attrName, attrVal) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res := traverse(c, attrName, attrVal)

		if res != nil {
			return res
		}
	}

	return nil
}

func GetElementByAttribute(n *html.Node, attrName string, attrVal string) *html.Node {
	return traverse(n, attrName, attrVal)
}
