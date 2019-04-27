//
// Pretty Path - print "path" env var on multiple lines
//
// just a test to play with go

package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// WIN_PATHS_LIST_SEP is separator char in paths list for windows
const WIN_PATHS_LIST_SEP = ";"
const WIN_PATH_SEP = "\\"
const NIX_PATHS_LIST_SEP = ":"
const NIX_PATH_SEP = "/"

func printListOfPaths(varValue, sep string) {
	paths := strings.Split(varValue, sep)
	max := len(paths)
	digits := int(math.Floor(math.Log10(float64(max)))) + 1
	decorator := "%" + strconv.Itoa(digits) + "d) "
	for i, p := range paths {
		fmt.Printf(decorator, i)
		fmt.Println(p)
	}
}

func main() {
	// cool, we can select OS separator at runtime
	listSeparator := NIX_PATHS_LIST_SEP
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		listSeparator = WIN_PATHS_LIST_SEP
	}

	for _, ev := range os.Environ() {
		v := strings.Split(ev, "=")
		if v[0] == "PATH" {
			printListOfPaths(v[1], listSeparator)
		}
	}
}
