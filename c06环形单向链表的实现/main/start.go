package main

import "fmt"

//环形单向链表的实现
func main() {
	head := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	Add(head, cat1)
	Add(head, cat2)
	Add(head, cat3)
	List(head)

	head = Del(head, 4)
	fmt.Println()
	List(head)
}

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func Add(head *CatNode, newNode *CatNode) {
	//如果添加的是第一个节点
	if head.next == nil {
		//那么添加节点就是头节点
		head.no = newNode.no
		head.name = newNode.name
		//指向头节点,形成环状
		head.next = head
		return
	}
	temp := head
	for {
		if temp.next == head {
			//遍历到最后
			break
		}
		temp = temp.next
	}
	//加入到链表的最后
	temp.next = newNode
	newNode.next = head
}

func List(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	//遍历输出
	for {
		fmt.Printf("节点:[id:%d name:%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func Del(head *CatNode, id int) *CatNode {
	//头指标
	temp := head
	//尾指标
	helper := head
	if temp.next == nil {
		fmt.Println("空链表")
		return head
	}
	//该链表只有一个节点的情况
	if temp.next == head {
		if temp.no == id {
			//如果要删除的节点是该链表唯一节点
			temp.next = nil
		}
		return head
	}
	//指向末尾节点
	for {
		if helper.next == head {
			//尾指标指向头节点,构成环形,即helper指向了末位
			break
		}
		helper = helper.next
	}
	//用于判定是否完成删除
	flag := true
	for {
		//temp指向了尾节点 [但还没进行最后节点的比较]
		if temp.next == head {
			break
		}
		//找到了要删除的节点
		if temp.no == id {
			//说明要删除的节点是头节点
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("删除节点:[%d]\n", id)
			flag = false
			break
		}
		//移动辅助指针
		temp = temp.next
		helper = helper.next
	}
	//补全逻辑 [进行最后节点的比较]
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("要删除的是 末位节点:[%d]\n", id)
		} else {
			fmt.Printf("未找到待删除的节点:[%d]\n", id)
		}
	}
	//最后完成对main栈head的修改(当前函数操作的是当前方法栈的head)
	return head
}
