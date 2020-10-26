package main

import "fmt"

//稀疏数组sparseArray 即二位数组进阶,压缩.只保存有用的数值,默认值(0)不做记录
func main() {
	//二维数组
	var chessArr [11][11]int
	chessArr[1][2] = 1 //第1行,第2列的值为1
	chessArr[2][3] = 2
	//输出二维数组
	for _, v := range chessArr {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//------------------------------------------------------------
	//转成稀疏数组
	var sparseArr []ValNode
	//因为go语言必须声明数组范围,所以第一node用于记录二维数组的范围
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	//将第一个node放入稀疏数组
	sparseArr = append(sparseArr, valNode)
	//遍历二位数组,转为稀疏数组
	for i, v := range chessArr {
		for j, v2 := range v {
			if v2 != 0 {
				//如果值不为0(默认值),则创建一个node节点进行记录
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				//添加到稀疏数组
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	//-------------------------------------------------
	//稀疏数组恢复成二维数组
	var chessArr2 [11][11]int
	for i, valNode := range sparseArr {
		//跳过第一条记录(记录数组范围的)
		if i != 0 {
			chessArr2[valNode.row][valNode.col] = valNode.val
		}
	}
	//输出恢复后的二位数组
	for _, v := range chessArr2 {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}
}

//定义node结构体记录二维数组的值
type ValNode struct {
	row int //行
	col int //列
	val int //值
}
