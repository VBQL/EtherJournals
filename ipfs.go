package main

/* TODO:
Pin files
Profit?
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"os"
)

func getInput() string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("File name to encode: ")
	text, _ := reader.ReadString('\n')
	// Cut out the newline character and return result
	return text[:len(text)-1]
}

func readFile(filename string) []byte{
	// Read data into @var rawBytes
	rawBytes, err:= ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return rawBytes
}

func writeToIPFS(alphaContent []byte){
	// Connect to localhost shell
	sh := shell.NewShell("localhost:5001")
	// @var id will be the QmHash
	id, err := sh.Add(bytes.NewReader(alphaContent))
	if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
	}
    fmt.Printf("IPFS hash: %s\n", id)
}

func main() {
	file := readFile(getInput())
	writeToIPFS(file)
}
