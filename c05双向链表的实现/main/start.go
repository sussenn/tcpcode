package main

import "fmt"

//双向链表 (相对单向链表, 双向链表可以从尾部到头部进行操作)
func main() {
	//头节点为空
	head := &HeroNode{}
	//待添加的新节点1
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}
	//添加到链表
	AddHeroNode(head, hero1)
	AddHeroNode(head, hero2)
	AddHeroNode(head, hero3)
	//显示
	ListHeroNode(head)
	//倒序打印
	fmt.Println()
	ListHeroNode2(head)
}

type HeroNode struct {
	no       int
	name     string
	nickname string
	//指向前一个节点
	pre *HeroNode
	//指向下一个节点
	next *HeroNode
}

//新增(无序)
func AddHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		//找到最后的节点,在尾部添加新节点
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//将新节点添加到链表后
	temp.next = newHeroNode
	//新节点的前部指向temp
	newHeroNode.pre = temp
}

//新增(排序)
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		//指向了链表最后
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			//当前指向节点的 下一个节点 比新插入的节点大, 即新插入的节点应在temp后
			break
		} else if temp.next.no == newHeroNode.no {
			//链表中存在相同序号的节点
			flag = false
			break
		}
		//辅助节点后移
		temp = temp.next
	}
	if !flag {
		fmt.Println("新增的节点已存在, no:", newHeroNode.no)
		return
	} else {
		//新节点在temp的下一个节点的前面
		newHeroNode.next = temp.next
		//新节点的头部指向temp
		newHeroNode.pre = temp
		//如果在最后的节点插入,则不执行这一步. 否则 temp.next.pre == nil
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		//再让 新节点成为temp的下一个节点
		temp.next = newHeroNode
	}
}

//显示双向链表信息
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	//遍历链表
	for {
		fmt.Printf("[%d %s %s]-->", temp.next.no, temp.next.name, temp.next.nickname)
		//辅助指针后移
		temp = temp.next
		//如果遍历到最后则跳出循环
		if temp.next == nil {
			break
		}
	}
}

//倒序输出双向链表
func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	//将temp定位到最后的节点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//遍历链表
	for {
		fmt.Printf("[%d %s %s]-->", temp.no, temp.name, temp.nickname)
		//辅助指针前移
		temp = temp.pre
		//如果遍历到头节点则跳出循环
		if temp.pre == nil {
			break
		}
	}
}

//删除
func DelHeroNode(head *HeroNode, no int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			//遍历到链表最后节点
			break
		} else if temp.next.no == no {
			//找到需要删除的节点
			flag = true
			break
		}
		//辅助指针后移
		temp = temp.next
	}
	if flag {
		//删除: 让temp后2位的节点顶替temp的后1位, 即抛弃了temp原先的后1位节点
		temp.next = temp.next.next
		//
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("删除的节点不存在")
	}
}
