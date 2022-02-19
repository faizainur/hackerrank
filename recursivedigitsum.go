package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigit(n string, k int32) int32 {
	// Write your code here
	// var p string

	// for i := 0; int32(i) < k; i++ {
	// 	p = p + n
	// }
	// pLen := len(p)
	// // pNum, err := strconv.ParseUint(p, 10, 128)
	// // if err != nil {
	// // 	checkError(err)
	// // }
	// var pNum big.Int
	// pNum.SetString(p, 10)

	// fmt.Println(p, pNum, pLen)

	// superDigit = recSum(pNum, int32(pLen), superDigit)
	// superDigitLen := len(strconv.Itoa(int(superDigit)))
	// superDigit = recSum(int64(superDigit), int32(superDigitLen), 0)
	// var superDigitLen int
	// firstIteration := true
	// for {
	// 	if firstIteration {
	// 		superDigit = recSum(pNum, int32(pLen), superDigit)
	// 		superDigitLen = len(strconv.Itoa(int(superDigit)))
	// 		firstIteration = false
	// 	} else {
	// 		superDigit = recSum(*big.NewInt(int64(superDigit)), int32(superDigitLen), 0)
	// 		superDigitLen = len(strconv.Itoa(int(superDigit)))
	// 	}

	// 	fmt.Println(superDigit, superDigitLen)

	// 	if superDigitLen == 1 {
	// 		break
	// 	}
	// }

	var result int64

	if len(n) == 1 {
		superDigit, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			checkError(err)
		}
		return int32(superDigit)
	}

	for _, v := range n {
		buf, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			checkError(err)
		}
		result += buf
	}

	return superDigit(strconv.FormatInt(result*int64(k), 10), 1)
}

func recSum(num big.Int, digits int32, sumTemp int32) int32 {
	var sum big.Int
	sum.SetInt64(int64(sumTemp))
	// pow := uint64(math.Pow(float64(10), float64(digits-1)))
	var pow big.Int
	// pow.SetUint64(uint64(math.Pow(float64(10), float64(digits-1))))
	pow.Exp(big.NewInt(10), big.NewInt(int64(digits-1)), nil)
	div := new(big.Int).Div(&num, &pow)
	rem := new(big.Int).Mod(&num, &pow)
	sum.Add(big.NewInt(sum.Int64()), div)
	fmt.Println(num.Int64(), digits, pow.Int64(), sum.Int64())
	if rem.Uint64() > 0 {
		return recSum(*rem, digits-1, int32(sum.Int64()))
	}
	return int32(sum.Int64())
}

func recSumS(num string) int32 {
	var sum int32

	for _, v := range []byte(num) {
		buf, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			checkError(err)
		}
		sum = sum + int32(buf)
	}

	fmt.Println(strconv.FormatUint(uint64(sum), 10))

	if len(strconv.FormatUint(uint64(sum), 10)) > 1 {
		return recSumS(strconv.FormatUint(uint64(sum), 10))
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
