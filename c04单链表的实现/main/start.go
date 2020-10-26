package main

import "fmt"

//结构体组成单链表
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
	//AddHeroNode(head, hero1)
	//AddHeroNode(head, hero2)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero3)
	//显示
	ListHeroNode(head)
	//删除
	fmt.Println()
	DelHeroNode(head, 1)
	ListHeroNode(head)
}

//表示一个节点
type HeroNode struct {
	//序号
	no       int
	name     string
	nickname string
	//指向下一个节点
	next *HeroNode
}

//新增(无序)
func AddHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//辅助指针,表示从头部节点开始
	temp := head
	for {
		//为空,表示找到了最后的节点,退出循环,在尾部添加新节点
		if temp.next == nil {
			break
		}
		//辅助节点往后移一位
		temp = temp.next
	}
	//将新节点添加到链表后
	temp.next = newHeroNode
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
		//先将 当前节点的 下一个节点 赋给新节点的 下一个节点. 即新节点在temp的下一个节点的前面
		newHeroNode.next = temp.next
		//再让 新节点成为temp的下一个节点
		temp.next = newHeroNode
	}
}

//显示链表信息	需要传入头节点
func ListHeroNode(head *HeroNode) {
	temp := head
	//如果没有下一个节点,则表示为空链表(head节点不存储元素)
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
	} else {
		fmt.Println("删除的节点不存在")
	}
}
