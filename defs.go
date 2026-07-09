package main

const (
	PUSH = 1
	ADD  = 2
	SUB  = 3
	MUL  = 4
	DIV  = 5
	OUT  = 6
	JMP  = 7
	JNZ  = 8
	JIZ  = 9
	POP  = 10
	DUP  = 11
)

var program = []int{
	PUSH, 5,
	PUSH, 3,
	ADD,
	OUT,
}

var stack []int
var pc int = 0
