// This file is part of GoIRS.
//
//    GoIRS is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    GoIRS is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with GoIRS.  If not, see <http://www.gnu.org/licenses/>.

package goirs

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/joffrey-bion/gosoup"
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
	if err != nil {
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

//Encontrar cuerpo (reescribir para cambiar el corpus)
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
