// Package insertanodeatthetailofalinkedlist provides solution for hackerrank problem: https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list
package insertanodeatthetailofalinkedlist

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

func (list *SinglyLinkedList) insert(data int32) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}

	list.tail = node
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
		llist.insert(llistItem)
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
