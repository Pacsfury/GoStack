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
	case COUT:
		fmt.Print(string(rune(stack[len(stack)-1])))
	case EQ:
		var result int = 0
		if (stack[len(stack)-1] == stack[len(stack)-2]) {
			result = 1
		} else {
			result = 0
		}
		stack = stack[:len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, result)
	case NEQ:
		var result int = 1
		if (stack[len(stack)-1] == stack[len(stack)-2]) {
			result = 0
		} else {
			result = 1
		}
		stack = stack[:len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, result)
	case LT:
		var result int = 1
		if (stack[len(stack)-1] < stack[len(stack)-2]) {
			result = 0
		} else {
			result = 1
		}
		stack = stack[:len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, result)
	case GT:
		var result int = 1
		if (stack[len(stack)-1] > stack[len(stack)-2]) {
			result = 0
		} else {
			result = 1
		}
		stack = stack[:len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, result)
	case LTE:
		var result int = 1
		if (stack[len(stack)-1] <= stack[len(stack)-2]) {
			result = 0
		} else {
			result = 1
		}
		stack = stack[:len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, result)
	case GTE:
		var result int = 1
		if (stack[len(stack)-1] >= stack[len(stack)-2]) {
			result = 0
		} else {
			result = 1
		}
		stack = stack[:len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, result)
	case AND:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a&b)
	case OR:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a|b)
	case XOR:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a^b)
	case NOT:
		if len(stack) < 1 {
			return
		}
		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, ^a)
	case SHL:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a<<b)
	case SHR:
		if len(stack) < 2 {
			return
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, a>>b)
	}
}