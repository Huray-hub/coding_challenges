// Package comparetwolinkedlists contains solution for hackerrank problem https://www.hackerrank.com/challenges/compare-two-linked-lists/problem
package comparetwolinkedlists

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type LinkedList struct {
	head *Node
}
type Node struct {
	data int
	next *Node
}

func (list *LinkedList) insert(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}

	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func compareLinkedLists(llistA, llistB LinkedList) int {
	currentA := llistA.head
	currentB := llistB.head

	for {
		if currentA.data != currentB.data {
			return 0
		}
		currentA = currentA.next
		currentB = currentB.next

		if currentA == nil {
			if currentB == nil {
				return 1
			} else {
				return 0
			}
		}

		if currentB == nil {
			return 0
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	testCases, err := strconv.Atoi(readLine(reader))
	checkError(err)

	results := make([]int, 0, testCases)

	var n, m int
	for i := 0; i < testCases; i++ {
		n, err = strconv.Atoi(readLine(reader))
		checkError(err)
		listA := prepareLinkedList(n, reader)

		m, err = strconv.Atoi(readLine(reader))
		checkError(err)
		listB := prepareLinkedList(m, reader)

		result := compareLinkedLists(listA, listB)
		results = append(results, result)
	}

	for _, result := range results {
		fmt.Fprintln(writer, result)
	}
	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func prepareLinkedList(nodesLen int, reader *bufio.Reader) LinkedList {
	list := LinkedList{}
	for i := 0; i < nodesLen; i++ {
		item, err := strconv.Atoi(readLine(reader))
		checkError(err)
		list.insert(item)
	}
	return list
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
