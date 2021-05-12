package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func insertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func shellSort(arr []int) {
	Arrlen := len(arr)
	var gap int
	for gap = Arrlen / 2; gap > 0; gap /= 2 {
		fmt.Printf("gap from begining is %d", gap)
		for i := gap; i < Arrlen; i++ {
			fmt.Printf("gap is %d", gap)
			fmt.Print("gap done")
			for j := i; j-gap >= 0; j = j - gap {
				fmt.Print("diff start")
				if arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
				}

			}
		}

	}
}

func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minindex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minindex] {
				minindex = j
			}
		}
		arr[i], arr[minindex] = arr[minindex], arr[i]
	}
}
