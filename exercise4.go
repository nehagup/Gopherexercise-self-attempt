package main

import (
	"flag"
	"golang.org/x/net/html"
	"strings"
	"io/ioutil"
	"fmt"
)

type Link struct{
	Href string
	Text string
}


func main(){
	file := flag.String("filename", "e4.html", "HTML file to be parsed")
	flag.Parse()
	b, _ := ioutil.ReadFile(*file)
	doc, _ := html.Parse(strings.NewReader(string(b)))
	Parse(doc)

}

func Parse(n *html.Node) ([]Link){
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println("\n" + a.Val)
			}
		}
		for a := n.FirstChild; a != nil; a = a.NextSibling {
			if a.Type == html.TextNode {
				fmt.Print(a.Data)
				continue
			}
			TagData(a)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Parse(c)
	}
	return nil
}

func TagData(n *html.Node)  {
	for a := n.FirstChild; a != nil; a = a.NextSibling {
		if a.Type == html.TextNode {
			fmt.Print(strings.TrimSpace(a.Data))
		}
	}
}