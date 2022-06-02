package SDM_lab2

import (
	"errors"
	"fmt"
)

type node struct {
	data     string
	previous *node
	next     *node
}

type doublyLinkedList struct {
	length int
	head   *node
	tail   *node
}

var errorEmpty = errors.New("empty list")
var errorIndex = errors.New("index out of range")
var errorNoNode = errors.New("node doesn't exist")

func initDList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (dList *doublyLinkedList) Length() int {
	return dList.length
}

func (dList *doublyLinkedList) Append(data string) {
	newNode := &node{
		data: data,
	}

	if dList.head == nil {
		dList.head = newNode
		dList.tail = newNode
	} else {
		currentNode := dList.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.previous = currentNode
		currentNode.next = newNode
		dList.tail = newNode
	}
	dList.length++
}

func (dList *doublyLinkedList) Insert(data string, index int) {
	newNode := &node{
		data: data,
	}

	newNode.next = nil
	newNode.previous = nil

	if dList.length == 1 || index > dList.length+1 {
		dList.Append(data)
	}

	if index < 1 {
		fmt.Println("Wrong index")
	} else {
		currentNode := dList.head
		for i := 1; i < index; i++ {
			if currentNode != nil {
				currentNode = currentNode.next
			}
		}
		if currentNode != nil {
			newNode.next = currentNode.next
			newNode.previous = currentNode
			currentNode.next = newNode
			if newNode.next != nil {
				newNode.next.previous = newNode
			}
		} else {
			fmt.Println(errorNoNode)
		}
	}
	dList.length++
}

func (dList *doublyLinkedList) Delete(index int) string {
	var delNode string
	var returnValue string

	if dList.head == nil {
		fmt.Println(errorEmpty)
		returnValue = ""
	}

	switch {
	case index < 0:
		fmt.Println(errorIndex)
		returnValue = ""

	case index == 0:
		dList.head = dList.head.next
		if dList.head != nil {
			dList.head.previous = nil
		} else {
			dList.tail = nil
		}
		delNode = dList.head.data
		dList.length--
		returnValue = delNode

	default:
		temp := dList.head
		n := 0
		for temp != nil && n < index {
			temp = temp.next
			n++
		}
		if temp == nil {
			fmt.Println(errorIndex)
		} else {
			temp.previous.next = temp.next
			if temp.next != nil {
				temp.next.previous = temp.previous
			} else {
				dList.tail = temp.previous
			}
			delNode = temp.data
			dList.length--
			returnValue = delNode
		}
	}
	return returnValue
}

func (dList *doublyLinkedList) DeleteAll(element string) {
	currentNode := dList.head
	for currentNode != nil {
		if currentNode.data == element {
			if dList.head == dList.tail {
				dList.head = nil
				dList.tail = nil
			} else if currentNode == dList.head {
				dList.head = currentNode.next
				dList.head.previous = nil
			} else if currentNode == dList.tail {
				dList.tail = currentNode.previous
				dList.tail.next = nil
			} else {
				currentNode.previous.next = currentNode.next
				currentNode.next.previous = currentNode.previous
			}
			dList.length--
		}
		currentNode = currentNode.next
	}

}

func (dList *doublyLinkedList) Get(index int) string {
	temp := dList.head
	pos := 0

	for pos != index && temp.next != nil {
		pos++
		temp = temp.next
	}

	if temp.data != "" {
		return temp.data
	}
	return ""
}

func (dList *doublyLinkedList) Clone() *doublyLinkedList {
	clone := initDList()
	temp := dList.head
	for i := 0; i < dList.length; i++ {
		clone.Append(temp.data)
		temp = temp.next
	}
	return clone
}

func (dList *doublyLinkedList) Reverse() {
	if dList.head == nil {
		fmt.Println(errorEmpty)
	}
	formerStart := dList.head
	formerAfter := formerStart.next
	formerStart.next = nil
	formerStart.previous = formerAfter
	for formerAfter != nil {
		formerAfter.previous = formerAfter.next
		formerAfter.next = formerStart
		formerStart = formerAfter
		formerAfter = formerAfter.previous
		dList.head = formerStart
	}
}

func (dList *doublyLinkedList) FindFirst(element string) int {
	temp := dList.head
	for i := 0; i < dList.length; i++ {
		if temp.data == element {
			return i
		}
		temp = temp.next
	}
	return -1
}

func (dList *doublyLinkedList) FindLast(element string) int {
	temp := dList.tail
	for i := dList.length - 1; i > 0; i-- {
		if temp.data == element {
			return i
		}
		temp = temp.previous
	}
	return -1
}

func (dList *doublyLinkedList) Clear() {
	for dList.head != nil {
		temp := dList.head
		dList.head = dList.head.next
		temp = nil
		fmt.Println(temp)
	}

}

func (dList *doublyLinkedList) Extend(anotherList *doublyLinkedList) {
	numNewElements := dList.Length()
	toRange := dList.head
	if toRange != nil {
		for numNewElements != 0 {
			anotherList.Append(toRange.data)
			toRange = toRange.next
			numNewElements--
		}
	} else {
		fmt.Println(errorEmpty)
	}
}

func (dList doublyLinkedList) displayData() {
	toPrint := dList.head
	lastElement := dList.tail
	if toPrint != nil {
		for dList.length != 1 {
			fmt.Printf("%v <=> ", toPrint.data)
			toPrint = toPrint.next
			dList.length--
		}
		fmt.Printf("%v \n", lastElement.data)
	} else {
		fmt.Println(errorEmpty)
	}
}
