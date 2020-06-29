package base

func BuildMaxHeap(a []int, heapSize int) {
	// 构建大顶堆
	for i := heapSize/2; i >= 0; i-- {
		MaxHeapify(a, i, heapSize)
	}
}

func MaxHeapify(a []int, i, heapSize int) {

	// 初始化左子节点、右子节点、父节点的角标
	l, r, largest := i * 2 + 1, i * 2 + 2, i

	if l < heapSize && a[l] > a[largest] {
		largest = l
	}

	if r < heapSize && a[r] > a[largest] {
		largest = r
	}

	// 如果父节点不是最大 则需要再次调整
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		MaxHeapify(a, largest, heapSize)
	}
}
