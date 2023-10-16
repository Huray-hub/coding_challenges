// Package findthemergepointoftwojoinedlinkedlists contains solution for hackerrank problem https://www.hackerrank.com/challenges/find-the-merge-point-of-two-joined-linked-lists/problem
package findthemergepointoftwojoinedlinkedlists

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type linkedList struct {
	head *node
}
type node struct {
	data int
	next *node
}

func (list *linkedList) insert(data int) {
	newNode := &node{data: data}
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

func findMergeNode(head1, head2 *node, len int) int {
	visited := make(map[*node]struct{}, len)
	current1 := head1
	current2 := head2

	for {
		if current1 != nil {
			if _, ok := visited[current1]; ok {
				return current1.data
			}
			visited[current1] = struct{}{}
			current1 = current1.next
		}
		if current2 != nil {
			if _, ok := visited[current2]; ok {
				return current2.data
			}
			visited[current2] = struct{}{}
			current2 = current2.next
		}
	}
}

func mergeLinkedLists(index int, head1, head2 *node, list1count, list2count int) {
	current1 := head1
	var i int
	for i < index && i < list1count {
		current1 = current1.next
		i++
	}
	current2 := head2

	for i = 0; i < list2count-1; i++ {
		current2 = current2.next
	}
	current2.next = current1
}

func printLinkedList(list linkedList, writer io.Writer) {
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

	var list1count, list2count, index int
	for i := 0; i < testCases; i++ {
		index, err = strconv.Atoi(readLine(reader))
		checkError(err)

		list1count, err = strconv.Atoi(readLine(reader))
		checkError(err)
		listA := prepareLinkedList(list1count, reader)

		list2count, err = strconv.Atoi(readLine(reader))
		checkError(err)
		listB := prepareLinkedList(list2count, reader)

		mergeLinkedLists(index, listA.head, listB.head, list1count, list2count)
		result := findMergeNode(listA.head, listB.head, list1count+list2count-index)

		fmt.Fprintf(writer, "%d\n", result)
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func prepareLinkedList(nodesLen int, reader *bufio.Reader) linkedList {
	list := linkedList{}
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
