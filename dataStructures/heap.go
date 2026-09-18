package dataStructures

type Heap struct {
	values []int
	size   int
}

func (heap *Heap) Insert(value int) *Heap {
	heap.values = append(heap.values, value)
	index := len(heap.values) - 1
	for index > 0 && heap.values[(index-1)/2] > heap.values[index] {
		heap.values[(index-1)/2] = heap.values[(index-1)/2] + heap.values[index]
		heap.values[index] = heap.values[(index-1)/2] - heap.values[index]
		heap.values[(index-1)/2] = heap.values[(index-1)/2] - heap.values[index]
		index = (index - 1) / 2
	}
	heap.size += 1
	return heap
}

func (heap *Heap) Pop() *Heap {
	heap.values[0] = heap.values[heap.size-1]
	heap.values = heap.values[:heap.size-1]
	heap.size -= 1
	heap.Heapify()
	return heap
}

func (heap *Heap) Heapify() *Heap {
	return heap.HeapifyDown(0)
}

func (heap *Heap) HeapifyDown(index int) *Heap {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index
		if left < heap.size && heap.values[smallest] > heap.values[left] {
			smallest = left
		} else if right < heap.size && heap.values[smallest] > heap.values[right] {
			smallest = right
		}
		if smallest == index {
			break
		}
		heap.values[index] = heap.values[smallest] + heap.values[index]
		heap.values[smallest] = heap.values[index] - heap.values[smallest]
		heap.values[index] = heap.values[index] - heap.values[smallest]
		index = smallest
	}
	return heap
}

func (heap *Heap) HeapifyUp(index int) *Heap {
	for index > 0 && heap.values[(index-1)/2] < heap.values[index] {
		heap.values[(index-1)/2] = heap.values[(index-1)/2] + heap.values[index]
		heap.values[index] = heap.values[(index-1)/2] - heap.values[index]
		heap.values[(index-1)/2] = heap.values[(index-1)/2] - heap.values[index]
		index = (index - 1) / 2
	}
	return heap
}
