# GoStack
A Stack-Based Virtual Machine built in Go.

---

## Introduction
GoStack is a virtual machine, but it's a little bit different from my other ones.

This is based on a stack, which is modified using some operations.

## Operations
| Operation | Value | Description |
|-----------|-------|-------------|
| PUSH      | 1     | Adds arg to stack |
| ADD       | 2     | Sums and deletes two last stack elements and appends result |
| SUB       | 3     | Substracts and deletes two last stack elements and appends result |
| MUL       | 4     | Multiplies and deletes two last stack elements and appends result |
| DIV       | 5     | Divides (if b!=0) and deletes two last stack elements and appends result |
| OUT       | 6     | Outputs the stack |
| JMP       | 7     | Jumps to arg |
| JNZ       | 8     | Jumps to arg if last element of stack != 0 |
| JIZ       | 9     | Jumps to arg if last element of stack == 0 |
| POP       | 10    | Deletes last element of stack |
| DUP       | 11    | Duplicates last element of stack |

## Error management
At the moment, there are two possible errors:
* Not enough args in stack
* b == 0 in division

Errors are treated as intentional, so they are just ignored.

## Example: sum 5 and 3
```go
var program = []int{
	PUSH, 5,
	PUSH, 3,
	ADD,
	OUT,
}
```
