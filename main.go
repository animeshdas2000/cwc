package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	cFlag := flag.Bool("c", false, "Used for counting number of bytes in the file")
	lFlag := flag.Bool("l", false, "Used for counting number of lines in the file")
	wFlag := flag.Bool("w", false, "Used for counting number of words in the file")
	mFlag := flag.Bool("m", false, "Used for counting Characters in the file")
	flag.Parse()
	path := flag.Arg(0)
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var out any
	switch {
	case *cFlag:
		out = calculateBytes(file)
	case *lFlag:
		out = calculateLines(file)
	case *wFlag:
		out = calculateWords(file)
	case *mFlag:
		out = calculateCharacters(file)
	default:
		out = fmt.Sprintf("%d\t%d\t%d\t%d", calculateBytes(file), calculateLines(file), calculateWords(file), calculateCharacters(file))

	}

	fmt.Printf("%v %s\n", out, path)
}

func calculateBytes(file []byte) int {
	//bytesSize := binary.Size(file)
	return len(file)
}

func calculateLines(file []byte) int {
	lines := 0
	for _, ch := range file {
		if ch == '\n' {
			lines++
		}
	}
	return lines
}

func calculateWords(file []byte) int {
	return len(strings.Fields(string(file)))
}

func calculateCharacters(file []byte) int {
	return utf8.RuneCount(file)
}
