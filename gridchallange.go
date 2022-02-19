package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
	// Write your code here
	if len(grid) == 1 {
		return "YES"
	}

	for i, v := range grid {

		buf := []byte(v)
		sort.Slice(buf, func(i, j int) bool {
			return buf[i] <= buf[j]
		})
		grid[i] = string(buf)

	}
	stringLen := len([]byte(grid[0]))
	// fmt.Println("Sorted : ", grid)

	for i := 0; i < stringLen; i++ {
		buf := []byte{}
		for j := 0; j < len(grid); j++ {
			buf = append(buf, []byte(grid[j])[i])
		}
		ok := checkIfSorted(buf)
		// fmt.Println(buf)
		// fmt.Println("Sorted ? ", ok)
		if !ok {
			return "NO"
		}
	}

	return "YES"

}

func checkIfSorted(arr []byte) bool {
	var res bool
	for i := 0; i < len(arr); i++ {
		if i != len(arr)-1 {
			if arr[i] <= arr[i+1] {
				res = true
			} else {
				res = false
				return res
			}
		}
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var grid []string

		for i := 0; i < int(n); i++ {
			gridItem := readLine(reader)
			grid = append(grid, gridItem)
		}

		result := gridChallenge(grid)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
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
