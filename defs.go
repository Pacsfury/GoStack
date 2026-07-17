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
	COUT = 12
)

var program = []int{
	1, 5,
	1, 3,
	2,
	6,
	1, 65,
	12,
}

var stack []int
var pc int = 0
