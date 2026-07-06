package main

import "fmt"

func execute(op int) {
	switch op {
	case PUSH:
		stack = append(stack, program[pc+1])
		pc += 1
	case ADD:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a+b)
	case SUB:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a-b)
	case MUL:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a*b)
	case DIV:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		if b == 0 {
			return
		}
		stack = stack[:len(stack)-2]
		stack = append(stack, a/b)
	case OUT:
		fmt.Println(stack)
	case JMP:
		pc = program[pc+1]
	case JNZ:
		if stack[len(stack)] != 0 {
			pc = program[pc+1]
		}
	case JIZ:
		if stack[len(stack)] == 0 {
			pc = program[pc+1]
		}
	case POP:
		stack = stack[:len(stack)-1]
	case DUP:
		stack = append(stack, stack[len(stack)])
	case RES:
		pc += program[pc+1]
	}
}