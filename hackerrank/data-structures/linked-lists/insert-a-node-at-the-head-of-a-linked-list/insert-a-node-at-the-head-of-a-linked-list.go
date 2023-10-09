// Package insertanodeattheheadofalinkedlist provides solution for hackerrank problem: https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list
package insertanodeattheheadofalinkedlist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int32
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (list *SinglyLinkedList) insertHead(data int32) {
	node := &Node{data: data, next: list.head}

	if list.head == nil {
		list.tail = node
	}

	list.head = node
}

func printLinkedList(head *Node, writer io.Writer) {
	currentHead := head
	for currentHead != nil {
		fmt.Fprintln(writer, currentHead.data)
		currentHead = currentHead.next
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist := SinglyLinkedList{}
	for i := 0; i < int(llistCount); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist.insertHead(llistItem)
	}

	printLinkedList(llist.head, writer)

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
