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
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var ptr1, ptr2 int32 = 0, 1
	var count int32 = 0

	fmt.Println("K Value = ", k)
	fmt.Println("N Value = ", n)
	fmt.Println("Array = ", ar)

	// for ptr1 < n {
	// 	for i, v := range ar[ptr2:n] {
	// 		if (ar[ptr1]+v)%k == 0 {
	// 			fmt.Println(ar[ptr1], v)
	// 			fmt.Println(ptr1, i)
	// 			count++
	// 		}
	// 	}

	// 	ptr1++
	// 	ptr2++
	// }

	for ptr1 < n && ptr2 < n {
		fmt.Println(ptr1, ptr2)
		if ptr1 < ptr2 && (ar[ptr1]+ar[ptr2])%k == 0 {
			count++
			fmt.Println(ptr2)
		}
		ptr2++

		if ptr2 == n && ptr1 < n {
			ptr1++
			ptr2 = ptr1 + 1
		}
	}

	return count

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := divisibleSumPairs(n, k, ar)

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
