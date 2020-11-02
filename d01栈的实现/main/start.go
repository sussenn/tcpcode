package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int //栈容量
	Top    int //栈顶	为-1表示栈底
	arr    [5]int
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	stack.List()
}

//入栈
func (this *Stack) Push(val int) (err error) {
	//先判断栈 是否已满
	if this.Top == this.MaxTop-1 {
		return errors.New("栈已满")
	}
	//栈底上移
	this.Top++
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	//先判断栈 是否为空
	if this.Top == -1 {
		return 0, errors.New("栈为空")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func (this *Stack) List() {
	//先判断栈 是否为空
	if this.Top == -1 {
		fmt.Println("栈为空")
		return
	}
	//从栈顶开始遍历
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}
