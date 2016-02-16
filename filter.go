package goirs

import (
	"bufio"
	"golang.org/x/net/html"
    "github.com/joffrey-bion/gosoup"
)


//Filter filtra el HTML de un documento proporcionado por input
func Filter(input *bufio.Reader, output *bufio.Writer) error {

	root, err := html.Parse(input)
    newroot := gosoup.WrapTree(root)
	if err != nil {
		return err
	}

	output.WriteString(extractTitle(root))
    output.WriteByte('\n')
    writeBody(newroot, output)
	return nil
}

func extractTitle(root *html.Node) string {
	var head *html.Node
	var title *html.Node

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "head" {
			head = c
			break
		}
	}

	if head == nil {
		return "No head"
	}

	for c := head.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "title" {
			title = c
			break
		}
	}

	if title == nil {
		return "No title"
	}

	if title.FirstChild.Type == html.TextNode {
		return title.FirstChild.Data
	}
	return "Wrong title"
}

//Encontrar el meollo del asunto
func writeBody(root *gosoup.Node, output *bufio.Writer) {
    //Vivan las funciones anónimas y la programación funcional!!!

    FindFirst(root, func(r *gosoup.Node) bool {
        return r.IsTag("div") && r.HasAttr("class") && r.Attr("class")=="node-inner"
    }).Descendants().Filter(func (n *gosoup.Node) bool {
        return n.IsTag("p")
    }).Apply(func (n *gosoup.Node){
        n.Descendants().Filter(func(m *gosoup.Node) bool {
            return m.Type == gosoup.TextNode
        }).Apply(func(m *gosoup.Node){
            output.WriteString(m.Data)
            output.WriteByte(' ')
        })
    })
}

//FindFirst está como retrocompatibilidad con el sistema anterior
func FindFirst(root *gosoup.Node, cond func(*gosoup.Node) bool) *gosoup.Node {
    iterator := root.Descendants().Filter(cond)
    return iterator.Next()
}
