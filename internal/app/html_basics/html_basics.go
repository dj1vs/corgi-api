package html_basics

import (
	"strings"

	"golang.org/x/net/html"
)

func GetAttribute(n *html.Node, key string) (string, bool) {

	for _, attr := range n.Attr {

		if attr.Key == key {
			// log.Println("->", attr.Val)
			return attr.Val, true
		}
	}

	return "", false
}

// children - how many levels deep should we go
func CollectText(n *html.Node, children int) string {
	text := ""
	if n.Type == html.TextNode {
		if checkIfTextIsValid(n.Data) {
			text += n.Data
			if n.Parent != nil && n.Parent.Data == "li" {
				text += "\n"
			}
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

func GetElementByData(n *html.Node, data string) *html.Node {
	// log.Println(n.Data)
	if n.Data == data {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		element := GetElementByData(c, data)
		if element != nil {
			return element
		}
	}

	return nil
}

func checkIfTextIsValid(text string) bool {
	if len(text) != 1 && strings.Count(text, "\n")+strings.Count(text, " ") == len(text) {
		return false
	} else if len(text) > 0 {
		return true
	}

	return false
}
