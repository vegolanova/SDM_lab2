# Lab 2: Unit tests, CI

*Order number in student list: 7 => 7 % 2 = 1 (Doubly Linked List)*

## Doubly Linked List
Doubly linked list is a linked data structure that consists of a set of sequentially linked records called nodes. Each node contains three fields: two link fields (references to the previous and to the next node in the sequence of nodes) and one data field. The beginning and ending nodes' previous and next links, respectively, point to some kind of terminator, typically a sentinel node or null, to facilitate traversal of the list. If there is only one sentinel node, then the list is circularly linked via the sentinel node. It can be conceptualized as two singly linked lists formed from the same data items, but in opposite sequential orders.

## Installation
1. Install Go.
2. Clone repository
```bash
  git clone https://github.com/vegolanova/SDM_lab2
```
3. Compile project from SDM_lab2 folder
4. Run tests
```
go run test
```

## Commit with failed test
https://github.com/vegolanova/SDM_lab2/commit/88453dc25260c62ce1536a51d107d521e96561f3
