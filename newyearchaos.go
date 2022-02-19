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
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
	// Write your code here

	counter := make(map[int32]int32)

	ptr, ptr2 := 0, 1
	for ptr < len(q) && ptr2 < len(q) {
		if q[ptr] > q[ptr2] {
			counter[q[ptr]]++
			ptr2++
		} else {
			ptr2++
		}

		if ptr2 == len(q) && ptr+1 < len(q) {
			ptr++
			ptr2 = ptr + 1
		}
	}

	// buf := make([]int32, len(q))
	// copy(buf, q)
	// sort.Slice(buf, func(i, j int) bool {
	// 	return buf[i] < buf[j]
	// })
	// fmt.Println(buf)

	// for i, v := range q {
	// 	diff := (v - 1) - int32(i)
	// 	if diff > 0 {
	// 		counter[v] = diff
	// 	}
	// }

	var res int32
	fmt.Println(counter)
	for _, v := range counter {
		if v <= 2 {
			res += v
		} else {
			fmt.Println("Too Chaotic")
			return
		}
	}
	fmt.Println(res)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
