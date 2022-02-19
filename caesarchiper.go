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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here

	stringBytes := []byte(s)
	result := []byte{}

	fmt.Println(stringBytes)

	for _, v := range stringBytes {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			buf := v + byte(k)
			fmt.Println("Not exceed ", buf)
			if v <= 90 && v >= 65 && buf > 90 {
				buf = getReminder(buf, 64, 90)
				fmt.Println("exceed ", buf)
				result = append(result, buf)

			} else if v <= 122 && v >= 97 && buf > 122 {
				buf = getReminder(buf, 96, 122)
				fmt.Println("exceed ", buf)
				result = append(result, buf)
			} else {
				result = append(result, buf)
			}
		} else {
			result = append(result, v)
		}
	}
	fmt.Println(result)
	resultString := string(result)
	return resultString
}

func getReminder(v, bot, top byte) byte {
	v = bot + (v - top)
	if v > top {
		v = getReminder(v, bot, top)
	}
	return v
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	fmt.Println(nTemp)
	result := caesarCipher(s, k)
	fmt.Println(result)

	fmt.Fprintf(writer, "%s\n", result)

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
