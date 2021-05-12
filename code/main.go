package main

import (
	"fmt"
)

func main() {
	tb := new(TbTree)
	tb.Key = 61
	for _, value := range []int{59, 87, 47, 60, 73, 98} {
		tb.Insert(value)
	}
	fmt.Printf("tb.key is %v\n", tb.Key)
	fmt.Printf("tb.LChild is %v\n", tb.LChild.Key)

	ptb, ntb, drt := tb.Search(60)

	if ptb != nil {
		fmt.Printf("parent of ntb is %v, with ntb is %v, the drt is %v", ptb.Key, ntb.Key, drt)
	}

	if ntb == nil {
		fmt.Println("ntb is nil")
	}
	tb.Delete(47)
	tbl := tb.LChild
	fmt.Printf("parent is %v, Rchild is %v", tbl.Key, tbl.RChild.Key)
	fmt.Printf("Lchild is %v", tbl.LChild.Key)
}
