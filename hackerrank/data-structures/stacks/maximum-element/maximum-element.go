// Package maximumelement contains solution for https://hackerrank.com/challenges/maximum-element
package maximumelement

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMax' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY operations as parameter.
 */

type command string

const (
	push command = "1"
	pop  command = "2"
	peek command = "3"
)

type stack struct {
	items []int32
}

func newStack() stack {
	items := make([]int32, 0)
	return stack{
		items: items,
	}
}

func (s *stack) push(number int32) {
	if number == 0 {
		return
	}
	s.items = append(s.items, number)
}

func (s *stack) pop() int32 {
	result := s.items[len(s.items)-1]
	if len(s.items) == 0 {
		return 0
	}
	s.items = s.items[:len(s.items)-1]
	return result
}

// definitely not a common stack method
func (s *stack) max() int32 {
	var maxNum int32
	for _, item := range s.items {
		if item > maxNum {
			maxNum = item
		}
	}
	return maxNum
}

func getMax(operations []string) []int32 {
	s := newStack()
	result := make([]int32, 0)
	// Write your code here
	for _, operation := range operations {
		fields := strings.Fields(operation)
		cmd := command(fields[0])
		switch cmd {
		case push:
			number, err := strconv.Atoi(fields[1])
			checkError(err)
			s.push(int32(number))
		case pop:
			_ = s.pop()
		case peek:
			result = append(result, s.max())
		}
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var ops []string

	for i := 0; i < int(n); i++ {
		opsItem := readLine(reader)
		ops = append(ops, opsItem)
	}

	res := getMax(ops)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
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
