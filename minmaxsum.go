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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var min, max int64

	for i := 0; i < len(arr); i++ {
		buffArr := make([]int32, len(arr))
		copy(buffArr, arr)
		removedItemArr := removeItem(buffArr, i)
		sum := sumArray(removedItemArr)
		if max == 0 && min == 0 {
			max = sum
			min = sum
		} else {
			if sum < min {
				min = sum
			} else if sum > max {
				max = sum
			}
		}
	}

	fmt.Printf("%d %d", min, max)
}

func sumArray(arr []int32) int64 {
	var sum int64 = 0
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	return sum
}

func removeItem(arr []int32, index int) []int32 {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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

