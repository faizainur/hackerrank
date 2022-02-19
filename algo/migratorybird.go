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
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// Write your code here

	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i] < arr[j]
	// })

	// var id int32
	// counter := make([]int, 6)
	// fmt.Println(counter)
	// fmt.Println("Array ", arr)
	// for ptr1 < int32(len(arr)) && ptr2 < int32(len(arr)) {
	// 	if id == 0 {
	// 		id = arr[ptr1]
	// 	}
	// 	counter[arr[ptr1]]++

	// 	if ptr2+1 < int32(len(arr)) && arr[ptr2+1] != arr[ptr1] {
	// 		fmt.Println("test")
	// 		// ptr2++
	// 		ptr1 = ptr2 + 1
	// 		if counter[arr[ptr2]] > counter[id] {
	// 			id = arr[ptr2]
	// 		}
	// 		ptr2++
	// 	}
	// 	ptr2++
	// }

	// for _, v := range arr {
	// if id == 0 {
	// 	id = int32(v)
	// }
	// counter[v]++

	// id = v

	// if int32(i)+1 < int32(len(arr)) && arr[int32(i)+1] != int32(v) {
	// 	fmt.Println("test")
	// 	// ptr2++
	// 	if counter[v] > counter[id] {
	// 		id = int32(v)
	// 	}
	// }
	// }
	// fmt.Println(counter)
	return int32(len(arr))
	// return id
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	fmt.Println(arrCount)

	fmt.Println(strings.TrimSpace(readLine(reader)))

	// arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	// arrTemp := strings.Split(readLine(reader), " ")

	// var arr []int32

	// for i := 0; int32(i) < int32(arrCount); i++ {
	// 	fmt.Println(i)
	// 	fmt.Println(len(arrTemp))
	// 	fmt.Println(arrTemp[len(arrTemp)-1])
	// 	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	// 	checkError(err)
	// 	arrItem := int32(arrItemTemp)
	// 	arr = append(arr, arrItem)
	// }

	// result := migratoryBirds(arr)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
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
