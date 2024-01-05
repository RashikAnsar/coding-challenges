package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RashikAnsar/coding-challenges/internal/ccwc"
)

func main() {
	lines := flag.Bool("l", false, "count lines")
	words := flag.Bool("w", false, "count words")
	byteSize := flag.Bool("c", false, "count bytes")
	chars := flag.Bool("m", false, "count characters")
	flag.Parse()

	fileName := flag.Arg(0)

	var (
		file *os.File
		err  error
	)

	if fileName == "" {
		file = os.Stdin
		fileName = file.Name()
		// TODO: remove this conditional after fixing seek for stdin data.
		if !(*lines || *words || *chars || *byteSize) {
			fmt.Println("pass some flag for stdin data")
			return
		}
	} else {
		file, err = os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	defer file.Close()

	ccwc.Run(file, *lines, *chars, *words, *byteSize, fileName)
}
