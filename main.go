package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Defining a boolean flag -b to count the number of total bytes
	bytes := flag.Bool("b", false, "Count bytes")
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words/lines/bytes
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {

	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if !countLines && countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a Counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc

}
