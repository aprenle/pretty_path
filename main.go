//
// Pretty Path - print "path" env var on multiple lines
//
// just a test to play with go

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const separator = ":"

func main() {
	for _, ev := range os.Environ() {
		v := strings.Split(ev, "=")
		if v[0] == "PATH" {
			paths := strings.Split(v[1], separator)
			max := len(paths)
			digits := int(math.Floor(math.Log10(float64(max)))) + 1
			decorator := "%" + strconv.Itoa(digits) + "d) "
			for i, p := range paths {
				fmt.Printf(decorator, i)
				fmt.Println(p)
			}
		}
	}
}
