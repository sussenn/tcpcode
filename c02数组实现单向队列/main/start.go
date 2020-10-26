package main

import (
	"errors"
	"fmt"
)

//数组实现单向队列. 即该队列无法复用(最后front和rear未重置)
type Queue struct {
	//队列长度
	maxSize int
	//go数组必须指定长度
	array [5]int
	//队首指标 初始假定为-1
	front int
	//队尾指标 初始假定为-1
	rear int
}

//添加
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 {
		return errors.New("队列已满")
	}
	//队尾+1
	this.rear++
	//在队尾 加入
	this.array[this.rear] = val
	return
}

//获取
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空. 队首和队尾相等,即空队列
	if this.rear == this.front {
		return -1, errors.New("队列为空")
	}
	//队首+1
	this.front++
	//将队首 输出
	val = this.array[this.front]
	return val, err
}

//显示队列
func (this *Queue) ShowQueue() {
	//从队首遍历到队尾. (front:不包含首位元素[所以要+1], rear:包含末位元素)
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("arr[%d] = %d\n", i, this.array[i])
	}
}

func main() {
	queue := Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	//添加 2次
	err := queue.AddQueue(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("添加队列成功")
	}
	err = queue.AddQueue(2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("添加队列成功")
	}
	//取出
	val, err := queue.GetQueue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("取出队列数值:", val)
	}
	//展示
	queue.ShowQueue()
}
