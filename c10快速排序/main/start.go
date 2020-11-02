package main

import (
	"fmt"
	"math/rand"
)

//快速排序
//随机取一位数做中轴,分为大小左右两边数组.
//递归调用方法本身,左右两边的数组再取中轴数,区分大小
//以此类推...排序
func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}

//p1:数组左下标 p2:数组右下标 p3:待排序的数组
func QuickSort(left int, right int, array *[5]int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	temp := 0
	for l < r {
		//控制从小到大
		for array[l] < pivot {
			l++
		}
		//控制从小到大
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	//避免死循环
	if l == r {
		l++
		r--
	}
	//递归调用
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}

}
