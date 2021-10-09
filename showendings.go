package main

import (
	"fmt"
	"io/ioutil"
)

func showTextFileEnding(file string) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	ending := bs[len(bs)-4:]
	fmt.Printf("|%3.2X |%3.2X  |%3.2X |%3.2X | %40.40s |\n",
		ending[0],
		ending[1],
		ending[2],
		ending[3],
		file)
}

func main() {

	showTextFileEnding("example01-utf8-lf.txt")
	showTextFileEnding("example01-utf8-crlf.txt")
	showTextFileEnding("example01-utf8bom-lf.txt")
	showTextFileEnding("example01-utf8bom-crlf.txt")
	showTextFileEnding("example01-utf16le-lf.txt")
	showTextFileEnding("example01-utf16le-crlf.txt")
	showTextFileEnding("example01-utf16be-lf.txt")
	showTextFileEnding("example01-utf16be-crlf.txt")

}
