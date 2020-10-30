package main

import (
	"fmt"
	"math/rand"
)

//插入排序
//创建一个有序数组,将待排序数组元素一次比较,并插入到有序数组中,最终完成排序
//原始数组:[23,0,12,56,34]
//默认第一位数为有序:
//[23] 0,12,56,34
//第1次找到插入位置:	[拿0跟有序数组里数值进行比较]
//[23,0] 12,56,34
//第2次找到插入位置:	[拿12跟有序数组里数值进行比较]
//[23,0,0] 56,34	//(12>0 所以 0 前一位复制出同大小数值0进行占位)
//[23,12,0] 56,34	//(12<23 找到位置)
//第3次找到插入位置:	[拿56跟有序数组里数值进行比较]
//[23,12,0,0] 34	//(56>0 复制0进行占位)
//[23,12,12,0] 34	//(56>12 复制12进行占位)
//[23,23,12,0] 34	//(56>23 复制23进行占位)
//[56,23,12,0] 34	//(最终找到插入位置)
//第四次找到插入位置: 	[拿34跟有序数组里数值进行比较]
//[56,23,12,0,0]	//(34>0 复制0进行占位)
//[56,23,12,12,0]	//(34>12 复制12进行占位)
//[56,23,23,12,0]	//(34>23 复制23进行占位)
//[56,34,23,12,0]	//(34<56 最终找到插入位置)

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(500)
	}
	InsertSort(&arr)
}

func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次排序插入arr:%v\n", i, *arr)
	}
	/*	//第一次排序,给第二个元素找到位置
		//表示插入数值本身
		insertVal := arr[1]
		//表示待插入位置下标
		insertIndex := 1 - 1
		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			//数据后移一位
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入	刚好选中最大数则可以少进行一次位置交换
		if insertIndex+1 != 1 {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("第1次排序插入:", *arr)

		//第2次排序,给第3个元素找到位置
		insertVal = arr[2]
		insertIndex = 2 - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != 2 {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("第2次排序插入:", *arr)*/
}
