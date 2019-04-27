//
// Pretty Path - print "path" env var on multiple lines
//
// it is just a test to experiment with golang
//
// usage:
//     pretty_print [options]
//
// options:
//   -g int
//         Lines group size; 0 to disable group (default 10)
//   -n    Print line number
//   -s    Sort paths before print
//
// examples:
//  $ pretty_print -s        //print sorted list of paths in PATH
//  $ pretty_print -n        //print list of paths in PATH with number prefix
//  $ pretty_print -n -g=4   //print paths in PATH with number prefix in groups of 4 lines
//

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

// WIN_PATHS_LIST_SEP is separator char in paths list for windows
const WIN_PATHS_LIST_SEP = ";"
const WIN_PATH_SEP = "\\"
const NIX_PATHS_LIST_SEP = ":"
const NIX_PATH_SEP = "/"

var printNumFlag bool
var linesGroupSize int
var sortFlag bool

var lineFormatter = "  * %s\n"

func printEntryWithNum(n int, p string) {
	fmt.Printf(lineFormatter, n+1, p)
}

func printEntryDot(n int, p string) {
	fmt.Printf(lineFormatter, p)
}

func printAllEntries(varValue, sep string, groupSize int) {
	paths := strings.Split(varValue, sep)
	sorted := ""
	if sortFlag {
		sort.Strings(paths)
		sorted = " (sorted)"
	}

	max := len(paths)
	fmt.Printf("\nPATH contains %d entries%s:\n", max, sorted)
	printFn := printEntryDot

	// Change print function if required
	if printNumFlag {
		digits := int(math.Floor(math.Log10(float64(max)))) + 1
		lineFormatter = "  %" + strconv.Itoa(digits) + "d) %s\n"
		printFn = printEntryWithNum
	}

	// actual print
	for i, p := range paths {
		if groupSize > 0 && (i%groupSize == 0) {
			fmt.Println()
		}
		printFn(i, p)
	}
}

func main() {
	// cool, we can select OS separator at runtime
	listSeparator := NIX_PATHS_LIST_SEP
	if runtime.GOOS == "windows" {
		listSeparator = WIN_PATHS_LIST_SEP
	}

	// parse options on command line
	flag.BoolVar(&printNumFlag, "n", false, "Print line number")
	flag.IntVar(&linesGroupSize, "g", 10, "Lines group size; 0 to disable group")
	flag.BoolVar(&sortFlag, "s", false, "Sort paths before print")
	flag.Parse()

	// grab PATH var and split value
	for _, ev := range os.Environ() {
		v := strings.Split(ev, "=")
		if v[0] == "PATH" {
			printAllEntries(v[1], listSeparator, linesGroupSize)
		}
	}
}
