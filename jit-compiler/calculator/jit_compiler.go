package calculator

import (
	"syscall"
)

type Compiler struct {
	mmapFunc   []byte
	currentIdx int
}

func NewCompiler() *Compiler {
	mmap, err := syscall.Mmap(
		-1,
		0,
		1024,
		syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC,
		syscall.MAP_PRIVATE|syscall.MAP_ANON,
	)
	if err != nil {
		panic(err)
	}
	return &Compiler{
		mmapFunc: mmap,
	}
}

func (c *Compiler) pushCode(code []uint8) {
	for _, b := range code {
		c.mmapFunc[c.currentIdx] = b
		c.currentIdx++
	}
}

func (c *Compiler) genCodeAST(ast Node) {
	switch node := ast.(type) {
	case *IntegerLiteral:
		c.pushCode([]uint8{0x6a, uint8(node.Value)}) // push {}
	case *InfixExpression:
		c.genCodeAST(node.Left)
		c.genCodeAST(node.Right)
		c.pushCode([]uint8{0x5f}) // pop rdi
		c.pushCode([]uint8{0x58}) // pop rax
		switch node.Operator {
		case "+":
			c.pushCode([]uint8{0x48, 0x01, 0xf8}) // add rax, rdi
		case "-":
			c.pushCode([]uint8{0x48, 0x29, 0xf8}) // sub rax, rdi
		case "*":
			c.pushCode([]uint8{0x48, 0x0f, 0xaf, 0xc7}) // imul rax, rdi
		case "/":
			c.pushCode([]uint8{0x48, 0x99})       // cqo
			c.pushCode([]uint8{0x48, 0xf7, 0xff}) // idiv rdi
		}
		c.pushCode([]uint8{0x50}) // push rax
	}
}

func (c *Compiler) GenCode(ast Node) {
	c.genCodeAST(ast)
	c.pushCode([]uint8{0x58}) // pop rax
	c.pushCode([]uint8{0xc3}) // ret
}
