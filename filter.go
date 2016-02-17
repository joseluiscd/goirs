package goirs

import (
	"bufio"
	"github.com/joffrey-bion/gosoup"
)

//Filter filtra el HTML de un documento proporcionado por input
func Filter(input *bufio.Reader, output *bufio.Writer) error {

	root, err := gosoup.Parse(input)
	if err != nil {
		return err
	}

	newroot := root.ChildrenByTag("html").Next()

	output.WriteString(extractTitle(newroot))
	output.WriteByte('\n')
	output.Flush()

	writeBody(newroot, output)
	output.Flush()
	return nil
}

func extractTitle(root *gosoup.Node) string {
	head := root.ChildrenByTag("head").Next()
	if head == nil {
		return "No head"
	}

	title := head.ChildrenByTag("title").Next()
	if title == nil {
		return "No title"
	}

	return title.FirstChild.Data
}

//Encontrar el meollo del asunto
func writeBody(root *gosoup.Node, output *bufio.Writer) {
	//Vivan las funciones anónimas y la programación funcional!!!

	FindFirst(root, func(r *gosoup.Node) bool {
		return (r.IsTag("div") && r.HasAttr("class") && r.Attr("class") == "node-inner")
	}).Descendants().Filter(func(n *gosoup.Node) bool {
		return n.IsTag("p")
	}).Apply(func(n *gosoup.Node) {
		n.Descendants().Filter(func(m *gosoup.Node) bool {
			return m.Type == gosoup.TextNode
		}).Apply(func(m *gosoup.Node) {
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
