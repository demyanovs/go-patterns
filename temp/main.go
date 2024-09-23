package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

const regex = `<.*?>`
const str = "<ul><li>test 1</li><li>test 2</li><li>test 3</li></ul><ul></ul><ul></ul>"

func main() {

	//r := strings.NewReader(str)
	//
	//doc, err := html.Parse(r)
	//if err != nil {
	//	panic("Fail to parse html")
	//}
	//x := traverse(doc)
	//fmt.Println(x)

	//fmt.Println(stripHtmlRegex(str))

	r := t2(str)
	fmt.Println(r)
}

func stripHtmlRegex(s string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, "")
}

func traverse(n *html.Node) string {
	if isUlElement(n) {
		return ul2rst(n)
	}

	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		rstText += traverse(c)
	}

	return rstText
}

func isUlElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ul"
}

func ul2rst(n *html.Node) string {
	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//if isLiElement(c) {
		for d := c.FirstChild; d != nil; d = d.NextSibling {
			fmt.Println(d.Data)
		} //}
	}

	return rstText
}

func t2(str string) string {

	values := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		return ""
	}
	table := doc.Find("ul")
	table.Find("li").Each(func(i int, li *goquery.Selection) {
		value := strings.TrimSpace(li.Text())
		value = fmt.Sprintf("%s", value)
		values = append(values, value)
	})

	return strings.Join(values, ", ")
}
