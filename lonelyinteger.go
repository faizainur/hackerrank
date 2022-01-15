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
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func lonelyinteger(a []int32) int32 {

    var lonelyInt int32
    // Write your code here

    if len(a) > 1 {

	    for i := 0; i < len(a); i++ {
		buffArr := make([]int32, len(a)) 
		copy(buffArr, a)
		removedItemArr := removeItem(buffArr, int32(i))
		for j := 0; j < len(removedItemArr); j++ {
			if removedItemArr[j] == a[i] {
				break;
			} else if removedItemArr[j] != a[i] && j == len(removedItemArr) - 1 {
				lonelyInt = a[i]
			}
		}

		if lonelyInt != 0 {
			break;
		}
	    }
    } else {
	    lonelyInt = a[0]
    }
    return lonelyInt

}

func removeItem(arr []int32, index int32) []int32 {
	arr[index] = arr[len(arr) -1]
	return arr[:len(arr) - 1]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    result := lonelyinteger(a)

    fmt.Fprintf(writer, "%d\n", result)

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

