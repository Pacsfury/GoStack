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

## Future Operations

### Memory & Pointers
* **LOAD** = 13: Reads data from a specific memory address and pushes it onto the data stack.
* **STORE** = 14: Pops a value from the data stack and writes it to a specific memory address.
* **LOADI** = 15: Indirect load. Pops an address from the stack, reads the value from that address, and pushes it.
* **STOREI** = 16: Indirect store. Pops an address and a value from the stack, then writes the value to that address.

### Comparisons
* **EQ** = 17: Pops two values and pushes 1 (true) if they are equal, otherwise 0 (false).
* **NEQ** = 18: Pops two values and pushes 1 if they are not equal, otherwise 0.
* **LT** = 19: Pops two values and pushes 1 if the first value is less than the second.
* **LTE** = 20: Pops two values and pushes 1 if the first value is less than or equal to the second.
* **GT** = 21: Pops two values and pushes 1 if the first value is greater than the second.
* **GTE** = 22: Pops two values and pushes 1 if the first value is greater than or equal to the second.

### Advanced Stack Manipulation
* **SWAP** = 23: Exchanges the positions of the top two items on the stack.
* **OVER** = 25: Copies the second item on the stack and pushes it to the top.
* **ROT** = 26: Rotates the top three items on the stack, moving the third item to the top.
* **PICK** = 27: Copies the n-th item down the stack (where n is popped from the top) and pushes it to the top.

### Functions & Return Stack
* **CALL** = 28: Pushes the current execution pointer to the return stack and jumps to a function address.
* **RET** = 29: Pops the return address from the return stack and jumps back to resume previous execution.
* **RUSH** = 30: Pops a value from the main data stack and pushes it directly onto the return stack.
* **RPOP** = 31: Pops a value from the return stack and pushes it onto the main data stack.

### Bitwise Operations
* **AND** = 32: Pops two values, performs a bitwise AND operation, and pushes the result.
* **OR** = 33: Pops two values, performs a bitwise OR operation, and pushes the result.
* **XOR** = 34: Pops two values, performs a bitwise exclusive OR operation, and pushes the result.
* **NOT** = 35: Pops the top value, inverts all its bits, and pushes the result.
* **SHL** = 36: Pops a value and a shift count, shifts the value bits left, and pushes the result.
* **SHR** = 37: Pops a value and a shift count, shifts the value bits right, and pushes the result.

### System & I/O
* **IN** = 38: Reads raw numerical data or hardware signal input and pushes it onto the stack.
* **CIN** = 39: Reads a single character from the standard input stream and pushes its character code.
* **HALT** = 40: Immediately stops all execution of the virtual machine program.
* **TIME** = 41: Pushes the current system clock time or tick count onto the stack.

### Math Extensions
* **MOD** = 42: Pops two values, calculates the remainder of their division, and pushes the result.
* **INC** = 43: Adds 1 directly to the top value on the stack.
* **DEC** = 44: Subtracts 1 directly from the top value on the stack.
