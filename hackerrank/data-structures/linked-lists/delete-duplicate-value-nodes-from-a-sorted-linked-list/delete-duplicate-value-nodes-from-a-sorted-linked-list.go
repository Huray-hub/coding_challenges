// Package deleteduplicatevaluenodesfromasortedlinkedlist contains solution for hackerrank problem https://www.hackerrank.com/challenges/delete-duplicate-value-nodes-from-a-sorted-linked-list/problem
package deleteduplicatevaluenodesfromasortedlinkedlist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprint(writer, sep)
		}
	}
}

/*
 * Complete the 'removeDuplicates' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts INTEGER_SINGLY_LINKED_LIST llist as parameter.
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func removeDuplicates(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	if llist == nil {
		return nil
	}
	currentNode := llist
	for currentNode.next != nil {
		if currentNode.data == currentNode.next.data {
			currentNode.next = currentNode.next.next
		} else {
			currentNode = currentNode.next
		}
	}
	return llist
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		llist := SinglyLinkedList{}
		for i := 0; i < int(llistCount); i++ {
			llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
			checkError(err)
			llistItem := int32(llistItemTemp)
			llist.insertNodeIntoSinglyLinkedList(llistItem)
		}

		llist1 := removeDuplicates(llist.head)

		printSinglyLinkedList(llist1, " ", writer)
		fmt.Fprintf(writer, "\n")
	}

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
