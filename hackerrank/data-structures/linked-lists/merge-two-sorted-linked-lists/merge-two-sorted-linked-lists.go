// Package mergetwosortedlinkedlists contains solution for hackerrank problem https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists/problem
package mergetwosortedlinkedlists

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

func mergeLinkedLists(listA, listB LinkedList) (mergedList LinkedList) {
	currentNodeA := listA.head
	currentNodeB := listB.head

	for currentNodeA != nil && currentNodeB != nil {
		if currentNodeA.data <= currentNodeB.data {
			mergedList.insert(currentNodeA.data)
			currentNodeA = currentNodeA.next
			continue
		}
		mergedList.insert(currentNodeB.data)
		currentNodeB = currentNodeB.next
	}

	for currentNodeA != nil {
		mergedList.insert(currentNodeA.data)
		currentNodeA = currentNodeA.next
	}
	for currentNodeB != nil {
		mergedList.insert(currentNodeB.data)
		currentNodeB = currentNodeB.next
	}

	return
}

func printLinkedList(list LinkedList, writer io.Writer) {
	currentNode := list.head
	for currentNode != nil {
		fmt.Fprintf(writer, "%d ", currentNode.data)
		currentNode = currentNode.next
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

	var n, m int
	for i := 0; i < testCases; i++ {
		n, err = strconv.Atoi(readLine(reader))
		checkError(err)
		listA := prepareLinkedList(n, reader)

		m, err = strconv.Atoi(readLine(reader))
		checkError(err)
		listB := prepareLinkedList(m, reader)

		mergedList := mergeLinkedLists(listA, listB)
		printLinkedList(mergedList, writer)
		fmt.Fprintf(writer, "\n")
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
