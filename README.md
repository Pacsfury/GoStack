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
| COUT      | 12    | Outputs to terminal the last stack item as ascii code |

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

## Future Operatioms

### Memory & Pointers
* **LOAD** = 13
* **STORE** = 14
* **LOADI** = 15
* **STOREI** = 16

### Comparisons
* **EQ** = 17
* **NEQ** = 18
* **LT** = 19
* **LTE** = 20
* **GT** = 21
* **GTE** = 22

### Advanced Stack Manipulation
* **SWAP** = 23
* **DROP** = 24
* **OVER** = 25
* **ROT** = 26
* **PICK** = 27

### Functions & Return Stack
* **CALL** = 28
* **RET** = 29
* **RUSH** = 30
* **RPOP** = 31

### Bitwise Operations
* **AND** = 32
* **OR** = 33
* **XOR** = 34
* **NOT** = 35
* **SHL** = 36
* **SHR** = 37

### System & I/O
* **IN** = 38
* **CIN** = 39
* **HALT** = 40
* **TIME** = 41

### Math Extensions
* **MOD** = 42
* **INC** = 43
* **DEC** = 44
