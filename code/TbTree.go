package main

import "fmt"

type TbTree struct {
	LChild *TbTree
	Key    int
	RChild *TbTree
}

func (tb *TbTree) Search(key int) (ptree, tree *TbTree, drt string) {

	if tb.Key == key {
		return nil, tb, ""
	}
	fmt.Printf("now search tbtree, key is %v\n", tb.Key)
	if tb.Key != key && tb.LChild == nil && tb.RChild == nil {
		fmt.Println("NO NODE FOUND")
		return nil, nil, ""
	}

	if tb.Key != key && tb.LChild != nil && tb.LChild.Key == key {
		fmt.Println("NODE FOUND LEFT")
		return tb, tb.LChild, "left"
	}

	if tb.Key != key && tb.RChild != nil && tb.RChild.Key == key {
		fmt.Println("NODE FOUND RIGHT")
		return tb, tb.RChild, "right"
	}

	if key > tb.Key {
		ptree, tree, drt = tb.RChild.Search(key)
	} else if key < tb.Key {
		ptree, tree, drt = tb.LChild.Search(key)
	}
	return ptree, tree, drt
}

func (tb *TbTree) Insert(key int) {
	if tb == nil {
		tb = new(TbTree)
		tb.Key = key
		return
	}

	if key > tb.Key {
		if tb.RChild == nil {
			nt := new(TbTree)
			nt.Key = key
			tb.RChild = nt
			fmt.Printf("Rchild insert done, key is %v\n", tb.RChild.Key)
			return
		}
		tb.RChild.Insert(key)
	} else if key < tb.Key {
		if tb.LChild == nil {
			nt := new(TbTree)
			nt.Key = key
			tb.LChild = nt
			fmt.Printf("Lchild insert done, key is %v\n", tb.LChild.Key)
			return
		}
		tb.LChild.Insert(key)
	}
	return
}

func (tb *TbTree) Delete(key int) {
	pdtb, dtb, drt := tb.Search(key)
	fmt.Printf("got dtb %v\n", dtb.Key)
	if dtb.LChild == nil && dtb.RChild == nil {
		switch drt {
		case "left":
			pdtb.LChild = nil
		case "right":
			pdtb.RChild = nil
		}

	} else if dtb.LChild != nil && dtb.RChild != nil {
		maxkey := dtb.GetMaxKeyInTree()
		fmt.Printf("max key is %v", maxkey)
		dtb.Delete(maxkey)
		dtb.Key = maxkey
	} else if dtb.LChild != nil && dtb.RChild == nil {
		pdtb.LChild = dtb.LChild
	} else if dtb.RChild != nil {
		pdtb.LChild = dtb.RChild
	}

}

func (tb *TbTree) GetMaxKeyInTree() (key int) {
	if tb == nil {
		return 0
	} else if tb.RChild == nil {
		return tb.Key
	}
	return tb.RChild.GetMaxKeyInTree()

}
