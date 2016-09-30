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

package main

// Programa que utiliza funciones de la librería GoIRS para limpiar y tokenizar
// texto de la entrada estándar.

import (
	"fmt"
	"os"

	"github.com/joseluiscd/goirs"
)

func main() {
	for x := range goirs.TokenizerIterator(os.Stdin) {
		fmt.Println(goirs.CleanToken(x))
	}
}
