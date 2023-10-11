// Package printinreverse contains solution for hackerrank problem https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list-in-reverse
package printinreverse

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

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string) {
	for node != nil {
		fmt.Printf("%d", node.data)

		node = node.next

		if node != nil {
			fmt.Print(sep)
		}
	}
}

/*
 * Complete the 'reversePrint' function below.
 *
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

type stack struct {
	first *stackNode
}
type stackNode struct {
	data int32
	next *stackNode
}

func (s *stack) push(data int32) {
	second := s.first
	s.first = &stackNode{
		data: data,
		next: second,
	}
}

func (s *stack) pop() int32 {
	result := s.first.data
	s.first = s.first.next
	return result
}

func (s *stack) empty() bool {
	return s.first == nil
}

func reversePrint(llist *SinglyLinkedListNode) {
	// Write your code here
	s := stack{}
	currentNode := llist
	for currentNode != nil {
		s.push(currentNode.data)
		currentNode = currentNode.next
	}

	for !s.empty() {
		fmt.Println(s.pop())
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	testsTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tests := int32(testsTemp)

	for testsItr := 0; testsItr < int(tests); testsItr++ {
		llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		llist := SinglyLinkedList{}
		for i := 0; i < int(llistCount); i++ {
			llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
			checkError(err)
			llistItem := int32(llistItemTemp)
			llist.insertNodeIntoSinglyLinkedList(llistItem)
		}

		reversePrint(llist.head)
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
