package main
import "fmt"

const (
	PUSH = 1
	ADD  = 2
	SUB  = 3
	MUL  = 4
	DIV  = 5
	OUT  = 6
	JMP  = 7
	JNZ  = 8
)

var program = []int{
	PUSH, 5,
	PUSH, 3,
	ADD,
	OUT,
}

var stack []int
var pc int = 0

func execute(op int) {
    switch(op) {
        case PUSH:
            stack = append(stack, program[pc+1])
            pc += 1
        case ADD:
            if len(stack)<2 {return}
            a := stack[len(stack)-1]
            b := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            stack = append(stack, a+b)
        case SUB:
            if len(stack)<2 {return}
            a := stack[len(stack)-1]
            b := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            stack = append(stack, a-b)
        case MUL:
            if len(stack)<2 {return}
            a := stack[len(stack)-1]
            b := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            stack = append(stack, a*b)
        case DIV:
            if len(stack)<2 {return}
            a := stack[len(stack)-1]
            b := stack[len(stack)-2]
            if b == 0 {return}
            stack = stack[:len(stack)-2]
            stack = append(stack, a/b)
        case OUT:
            fmt.Println(stack)
        case JMP:
            pc = program[pc+1]
        case JNZ:
            if stack[len(stack)] == 0 {pc = program[pc+1]}
            
   } 
}

func main() {

    for pc < len(program) {
        execute(program[pc])
        pc += 1
    }
}
