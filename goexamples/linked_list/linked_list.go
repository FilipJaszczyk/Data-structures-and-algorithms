package linkedlist

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func (p Person) Equal(given Person) bool {
	return p.ID == given.ID
}

type Node struct {
	Data     Person
	NextNode *Node
}

type LinkedList struct {
	Head          *Node
	NumberOfNodes int
}

func (l *LinkedList) InsertStart(data Person) {
	newNode := Node{Data: data}
	if l.Head != nil {
		newNode.NextNode = l.Head
	}
	l.Head = &newNode
	l.NumberOfNodes += 1
}

func (l *LinkedList) InsertEnd(data Person) {
	newNode := Node{Data: data}
	if l.Head == nil {
		l.Head = &newNode
		return
	}
	currentNode := l.Head
	for currentNode != nil {
		currentNode = currentNode.NextNode
	}
	currentNode.NextNode = &newNode
	l.NumberOfNodes += 1
}

func (l *LinkedList) Remove(data Person) {
	if l.Head == nil {
		return
	}
	if l.Head.Data.Equal(data) {
		l.Head = l.Head.NextNode
		return
	}
	currentNode := l.Head
	previousNode := &Node{}
	for currentNode != nil && !currentNode.Data.Equal(data) {
		previousNode = currentNode
		currentNode = currentNode.NextNode
	}
	if currentNode != nil {
		previousNode.NextNode = currentNode.NextNode
		l.NumberOfNodes -= 1
	}
}

func (l LinkedList) String() string {
	var b strings.Builder
	currentNode := l.Head
	for currentNode != nil {
		fmt.Fprint(&b, fmt.Sprint(currentNode.Data))
		currentNode = currentNode.NextNode
	}
	return b.String()
}
