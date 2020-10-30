package main

import (
	"fmt"
	"math/rand"
	"time"
)

//冒泡排序 < 选择排序 < 插入排序 < 快速排序
//选择排序:
//1.从所有数中选取最小数,拿最小数跟第一位置换
//2.再从第二位开始的所有数选取最小数,跟第二位置换
//3.再从第三位开始...
//4.以此类推排序....
func main() {
	var arr [100000]int
	for i := 0; i < 100000; i++ {
		arr[i] = rand.Intn(300000)
	}
	start := time.Now().Unix()
	Sort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时:%d秒\n", end-start)
	//fmt.Println(arr)
}

func Sort(arr *[100000]int) {
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}

	/*//思路代码
	//假设max为最大数
	max := arr[0]
	maxIndex := 0
	//遍历0位之后的数,找到真正的最大数
	for i := 0 + 1; i < len(arr); i++ {
		//max假定数小于真正的最大数,即找到真正最大数
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	//如果max刚好是真正的最大数,那可以少交换一次
	if maxIndex != 0 {
		//交换	5,100 = 100,5 [go可以直接交换位置!]
		arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
	}
	fmt.Println("第1次交换:", *arr)

	//第二次交换...
	max := arr[1]
	maxIndex := 1
	//遍历1位之后的数,找到真正的最大数
	for i := 1 + 1; i < len(arr); i++ {
		//max假定数小于真正的最大数,即找到真正最大数
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	//如果max刚好是真正的最大数,那可以少交换一次
	if maxIndex != 1 {
		//交换
		arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
	}
	fmt.Println("第2次交换:", *arr)

	*/
}
