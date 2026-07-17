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
	EQ   = 13
	NEQ  = 14
	LT   = 15
	GT   = 16
	LTE  = 17
	GTE  = 18
	AND  = 19
	OR   = 20
	XOR  = 21
	NOT  = 22
	SHL  = 23
	SHR  = 24
)

var program = []int{
	PUSH, 5,
	PUSH, 90,
	SUB,
	OUT,
}

var stack []int
var pc int = 0
