package main

import "fmt"

//环形链表解决约瑟夫问题
//约瑟夫问题: n人围成一圈,从k开始报数,数到m(k~m),m出列,m+1开始继续报数到m... 直至所有人出列,得到依次出列人的序号队列
func main() {
	first := Add(500)
	//1
	fmt.Println("头头:", first.No)
	Show(first)
	Joseph(first, 20, 31)
}

type Boy struct {
	No   int
	Next *Boy
}

//添加num个节点到链表,返回头节点
func Add(num int) *Boy {
	//头节点
	first := &Boy{}
	//辅助节点
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("传参必须>1")
		return first
	}
	//循环构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//先确定头节点
		if i == 1 {
			first = boy
			curBoy = boy
			//指向自己,构成环形
			curBoy.Next = first
		} else {
			//指向下一个节点,构建链表
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

func Show(first *Boy) {
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	//辅助指针
	curBoy := first
	for {
		fmt.Printf("节点编号:%d ->", curBoy.No)
		//末位指向首位,即完成遍历
		if curBoy.Next == first {
			break
		}
		//辅助指针后移
		curBoy = curBoy.Next
	}
}

//p: 头节点/起始数(k)/报数(m)
func Joseph(first *Boy, startNo int, num int) {
	if first == nil {
		fmt.Println("空链表")
		return
	}
	//辅助指针,尾随first(first作遍历指针,其指向谁则谁出列)
	tail := first
	//先将tail指向最后节点
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	//让first移动到起始数
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	for {
		for i := 1; i <= num-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("出列节点:%d\n", first.No)
		//first指向的节点出列
		first = first.Next
		tail.Next = first
		//如果最后只剩下一个节点
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出列节点:%d\n", first.No)
}
