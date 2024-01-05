package main

import (
	"fmt"
	"os"

	"github.com/RashikAnsar/coding-challenges/internal/ccwc"
)

func main() {
	var fileName string
	var inputFlag string

	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	if len(os.Args) > 2 {
		fmt.Println("condition true")
		fileName = os.Args[2]
		inputFlag = os.Args[1]
	} else {
		fmt.Println("condition false")
		fileName = os.Args[1]
	}

	file, err := os.Open(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s doesn't exist", fileName)
		} else {
			fmt.Println(err)
		}
		return
	}

	defer file.Close()

	switch inputFlag {
	case "-c":
		byteSize, err := ccwc.GetBytes(file)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(byteSize)
	case "-l":
		fmt.Println("Count lines")
		lines := ccwc.GetLines(file)
		fmt.Println(lines)
	case "-w":
		fmt.Println("count words")
		words := ccwc.GetWords(file)
		fmt.Println(words)
	case "-m":
		fmt.Println("count characters")
		chars := ccwc.GetChars(file)
		fmt.Println(chars)
	default:
		ccwc.DefaultOutput(file)
	}

}
