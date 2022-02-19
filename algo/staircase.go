package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	// fmt.Println(n)
	for i := 0; int32(i) < n; i++ {
		spaces := ""
		symbols := ""
		for j := 0; int32(j) < n-(int32(i)+1); j++ {
			spaces += " "
			// fmt.Println("asd" + "asd")
			// fmt.Printf("%s", spaces)
		}
		for k := 0; int32(k) < n-(n-(int32(i)+1)); k++ {
			symbols += "#"
			// fmt.Println("hell" + "asd")
			// fmt.Printf("%s", symbols)
		}
		fmt.Printf("%s%s", spaces, symbols)
		if int32(i)+1 < n {
			fmt.Println()
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
