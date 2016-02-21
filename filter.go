package goirs

import (
	"io"
	"github.com/joffrey-bion/gosoup"
	"os"
	"bufio"
	"bytes"
	"strings"
)

//Filter filtra el HTML de un documento proporcionado por input
func Filter(input io.Reader, output io.Writer) error {

	root, err := gosoup.Parse(input)
	if err != nil {
		return err
	}

	newroot := root.ChildrenByTag("html").Next()

	io.WriteString(output, extractTitle(newroot))
	io.WriteString(output, "\n")

	writeBody(newroot, output)
	return nil
}

//FilterFile abre el archivo especificado y lo filtra
func FilterFile(input string) io.Reader {
	file, err := os.Open(input)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	readfile := bufio.NewReader(file)
	buffer := bytes.NewBuffer(nil)

	Filter(readfile, buffer)

	return buffer
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

	return strings.TrimSpace(strings.Split(title.FirstChild.Data, "|")[0])
}

//Encontrar el meollo del asunto
func writeBody(root *gosoup.Node, output io.Writer) {
	//Vivan las funciones anónimas y la programación funcional!!!

	findFirst(root, func(r *gosoup.Node) bool {
		return (r.IsTag("div") && r.HasAttr("class") && r.Attr("class") == "node-inner")
	}).Descendants().Filter(func(n *gosoup.Node) bool {
		return n.IsTag("p")
	}).Apply(func(n *gosoup.Node) {
		n.Descendants().Filter(func(m *gosoup.Node) bool {
			return m.Type == gosoup.TextNode
		}).Apply(func(m *gosoup.Node) {
			io.WriteString(output, m.Data)
			io.WriteString(output, " ")
		})
	})
}

func findFirst(root *gosoup.Node, cond func(*gosoup.Node) bool) *gosoup.Node {
	iterator := root.Descendants().Filter(cond)
	return iterator.Next()
}
