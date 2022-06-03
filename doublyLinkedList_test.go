package main

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("C")

	got := testDLList.Length()
	want := 4

	if got != want {
		t.Errorf("Failed TestLength(): got %q, wanted %q", got, want)
	}
}

func TestAppend(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")

	got := testDLList.Length()
	want := 1

	if got != want {
		t.Errorf("Failed TestAppend(): got %q, wanted %q", got, want)
	}
}

func TestInsert(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")

	testDLList.Insert("B", 1)

	got := testDLList.Get(1)
	want := "B"

	if got != want {
		t.Errorf("Failed TestInsert(): got %q, wanted %q", got, want)
	}
}

func TestDelete(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("B")
	testDLList.Append("C")
	testDLList.Delete(1)

	got := testDLList.Length()
	want := 2

	if got != want {
		t.Errorf("Failed TestDelete(): got %q, wanted %q", got, want)
	}
}
func TestDelAll(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("A")
	testDLList.Append("C")

	testDLList.DeleteAll("C")

	c1 := testDLList.Get(1)
	c2 := testDLList.Get(2)
	c3 := testDLList.Get(5)

	allCs := []string{c1, c2, c3}
	prescence := []bool{}
	var tableFalse int

	for i := range allCs {
		prescence = append(prescence, allCs[i] == "C")
	}

	for j := 0; j < len(prescence); j++ {
		if prescence[j] {
			tableFalse = 0
		} else {
			tableFalse = 1
		}
	}

	got := tableFalse
	want := 1

	if got != want {
		t.Errorf("Failed TestDelAll(): got %v, wanted %v", got, want)

	}
}

func TestGet(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("A")
	testDLList.Append("D")

	valueC := testDLList.Get(1)

	got := valueC
	want := "C"

	if got != want {
		t.Errorf("Failed TestGet(): got %q, wanted %q", got, want)
	}
}

func TestClone(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("C")

	testDLListClone := testDLList.Clone()

	var same int

	if testDLList == testDLListClone {
		same = 1
	} else {
		same = 0
	}

	got := same
	want := 0

	if got != want {
		t.Errorf("Failed TestClone(): got %q, wanted %q", got, want)
	}
}

func TestReverse(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("C")
	testDLList.Append("C")
	testDLList.Append("D")

	initialListVals := []string{
		testDLList.head.data,
		testDLList.head.next.data,
		testDLList.head.next.next.data,
		testDLList.tail.previous.previous.data,
		testDLList.tail.previous.data,
		testDLList.tail.data,
	}

	fmt.Println(initialListVals)
	testDLList.displayData()

	testDLList.Reverse()
	testDLList.displayData()

	reversedListVals := []string{
		testDLList.Get(0),
		testDLList.Get(1),
		testDLList.Get(2),
		testDLList.Get(3),
		testDLList.Get(4),
		testDLList.Get(5),
	}
	fmt.Println(reversedListVals)

	var boolTable []bool
	var same int

	for i := 0; i < testDLList.length; i++ {
		boolTable = append(boolTable, initialListVals[i] == reversedListVals[i])
	}

	if boolTable[0] == true && boolTable[1] == true &&
		boolTable[2] == true && boolTable[3] == true &&
		boolTable[4] == true && boolTable[5] == true {
		same = 1
	} else {
		same = 0
	}

	got := same
	want := 0

	if got != want {
		t.Errorf("Failed TestReverse(): got %q, wanted %q", got, want)
	}
}

func TestFindFirst(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("C")

	firstC := testDLList.FindFirst("C")

	got := firstC
	want := 1

	if got != want {
		t.Errorf("Failed TestFindFirst(): got %q, wanted %q", got, want)
	}
}

func TestFindLast(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("C")

	lastC := testDLList.FindLast("C")

	got := lastC
	want := 3

	if got != want {
		t.Errorf("Failed TestFindLast(): got %q, wanted %q", got, want)
	}
}

func TestClear(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("C")

	testDLList.Clear()

	testDLList.displayData()
	fmt.Println(testDLList.length)

	got := testDLList.length
	want := 0

	if got != want {
		t.Errorf("Failed TestClear(): got %q, wanted %q", got, want)
	}
}

func TestExtend(t *testing.T) {
	testDLList := initDList()
	testDLList.Append("A")
	testDLList.Append("C")
	testDLList.Append("B")
	testDLList.Append("C")

	testDLList2 := initDList()
	testDLList2.Append("N")
	testDLList2.Append("M")
	testDLList2.Append("K")

	testDLList.Extend(testDLList2)
	got := testDLList2.length
	want := 7

	if got != want {
		t.Errorf("Failed TestExtend(): got %q, wanted %q", got, want)
	}
}
