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
	"strings"
	"testing"
)

func TestTokenizer(t *testing.T) {
	str := "Esté és el ñtexto.que 10.5  Ño  .debemos déé. . limpiáŕ"
	read := bufio.NewReader(strings.NewReader(str))
	it := TokenizerIterator(read)
	<-it
}
