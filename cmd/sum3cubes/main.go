/*
	sum3cubes calculates the sum of the cubes of 3 integers.
    Copyright (C) 2019 Esther Sánchez Aguilar.

	There is a copy of the GNU General Public License in the LICENSE
	file at the root of this repository. If not, see
	<https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log"

	"github.com/thersanchez/sum3cubes/pkg/sum3cubes"
)

func main() {
	results, err := sum3cubes.Do(2)
	if err != nil {
		log.Fatal(err)
	}
	for i, r := range results {
		// TODO: calculate indentation based on biggest number
		fmt.Printf("%5d: %v\n", i, r)
	}
}
