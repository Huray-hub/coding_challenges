// Package dynamicarray contains solution for hackerrank problem: https://www.hackerrank.com/challenges/dynamic-array
package dynamicarray

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	arr := make([][]int32, n)
	for i := int32(0); i < n; i++ {
		arr[i] = make([]int32, 0)
	}

	answers := make([]int32, 0, n)

	var x, y,
		idx,
		lastAnswer int32

	for _, query := range queries {
		x, y = query[1], query[2]
		idx = (x ^ lastAnswer) % n

		switch query[0] {
		case 1:
			arr[idx] = append(arr[idx], y)
		case 2:
			lastAnswer = arr[idx][y%int32(len(arr[idx]))]
			answers = append(answers, lastAnswer)
		}
	}
	return answers
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	RunTestCases, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer RunTestCases.Close()

	writer := bufio.NewWriterSize(RunTestCases, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := dynamicArray(n, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
