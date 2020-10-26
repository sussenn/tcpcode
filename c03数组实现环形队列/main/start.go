package main

import (
	"errors"
	"fmt"
	"os"
)

//环形队列的数组实现
type CircleQueue struct {
	//5 	队列的真实容量实际为4,预留1位给tail
	maxSize int
	//底层数组的大小为5
	array [5]int
	//指向队首	初始化为0	实际指向的位数即队首元素所在位
	head int
	//指向队尾	初始化为0	实际指向的位数即队尾元素所在位+1
	tail int
}

//判断队列是否已满
func (this *CircleQueue) IsFull() bool {
	//已满时: (4+1)%5 = 0
	return (this.tail+1)%this.maxSize == this.head
}

//判断队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	//头部为0, 尾部为1 即有1位元素
	return this.tail == this.head
}

//计算队列当前元素个数
func (this *CircleQueue) Size() int {
	//(5+5-0) %5 = 0 [tail指向预留位,即开始重置数组]
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

//新增
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	//在尾部新增
	this.array[this.tail] = val
	//(2+1) % 5 = 3 (尾部递增: tail从2递增到3)
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//取出
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列为空")
	}
	//取出头部
	val = this.array[this.head]
	//头部递增
	this.head = (this.head + 1) % this.maxSize
	return
}

//显示队列
func (this *CircleQueue) List() {
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//辅助变量,指向头部 	?????????
	temp := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", temp, this.array[temp])
		temp = (temp + 1) % this.maxSize
	}
	fmt.Println()
}

func main() {
	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入入队列数值:")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数:", val)
			}
		case "show":
			queue.List()
		case "exit":
			os.Exit(0)
		}
	}
}
