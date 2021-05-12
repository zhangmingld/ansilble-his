package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type AVLTree struct {
	key    int
	LChild *AVLTree
	RChild *AVLTree
	height int
}

func (avl *AVLTree) lspin() {
	al := avl.LChild
	avl.LChild = al.RChild
	al.RChild = avl
	avl.height = max(avl.LChild.height, avl.RChild.height) + 1
	al.height = max(al.LChild.height, al.RChild.height) + 1
	avl = al
}

func (avl *AVLTree) rspin() {
	ar := avl.RChild
	avl.RChild = ar.LChild
	ar.LChild = avl
	avl.height = max(avl.LChild.height, avl.RChild.height) + 1
	ar.height = max(ar.LChild.height, ar.RChild.height) + 1
}

func (avl *AVLTree) ldspin() {
	al := avl.LChild
	alr := al.RChild
	avl.LChild = alr
	al.RChild = alr.LChild
	alr.LChild = al
	al.height = max(al.LChild.height, al.RChild.height) + 1
	alr.height = max(alr.LChild.height, alr.RChild.height) + 1
	avl.height = max(avl.LChild.height, avl.RChild.height) + 1
	avl.lspin()

}

func (avl *AVLTree) rdspin() {
	ar := avl.RChild
	arl := ar.LChild
	avl.RChild = arl
	ar.LChild = arl.RChild
	arl.RChild = ar
	ar.height = max(ar.LChild.height, ar.RChild.height) + 1
	arl.height = max(arl.LChild.height, arl.RChild.height) + 1
	avl.height = max(avl.LChild.height, avl.RChild.height) + 1
	avl.lspin()

}

func (avl *AVLTree) GetMaxValue() int {
	if avl.RChild == nil {
		return avl.key
	}
	return avl.RChild.GetMaxValue()
}

func (avl *AVLTree) GetMinValue() int {
	if avl.LChild == nil {
		return avl.key
	}
	return avl.LChild.GetMinValue()
}

func (avl *AVLTree) Insert(key int) {
	if key > avl.key {
		if avl.RChild == nil {
			navl := new(AVLTree)
			avl.RChild = navl
		} else {
			avl.RChild.Insert(key)
			if avl.LChild.height-avl.RChild.height > 1 {
				if key > avl.RChild.key {
					avl.rspin()
				} else {
					avl.rdspin()
				}
			}
		}
	}

	if key < avl.key {
		if key > avl.key {
			if avl.RChild == nil {
				navl := new(AVLTree)
				avl.RChild = navl
			} else {
				avl.LChild.Insert(key)
				if avl.LChild.height-avl.RChild.height > 1 {
					if key < avl.LChild.key {
						avl.lspin()
					} else {
						avl.ldspin()
					}
				}
			}
		}
	}

	avl.height = max(avl.LChild.height, avl.RChild.height) + 1
}

func (avl *AVLTree) Delete(key int) *AVLTree {
	if avl.key == key {
		if avl.LChild != nil && avl.RChild != nil {
			if avl.LChild.height > avl.RChild.height {
				maxKey := avl.LChild.GetMaxValue()
				avl.key = maxKey
				avl.LChild = avl.LChild.Delete(maxKey)
			} else {
				minKey := avl.RChild.GetMinValue()
				avl.key = minKey
				avl.RChild = avl.RChild.Delete(minKey)
			}
		} else if avl.LChild != nil {
			avl = avl.LChild
		} else {
			avl = avl.RChild
		}
	} else if avl.key < key {
		avl.RChild = avl.RChild.Delete(key)
		if avl.RChild.height-avl.LChild.height > 1 {
			avl.rspin()
		} else if avl.RChild.height-avl.LChild.height < -1 {
			avl.rdspin()
		}
	} else {
		avl.LChild = avl.LChild.Delete(key)
		if avl.LChild.height-avl.RChild.height > 1 {
			avl.lspin()
		} else if avl.RChild.height-avl.LChild.height < -1 {
			avl.ldspin()
		}
	}
	avl.height = max(avl.LChild.height, avl.RChild.height) + 1

	return avl
}
