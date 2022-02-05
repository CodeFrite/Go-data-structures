package heap

import (
	"errors"
)

// The element data[0] is a dummy elt that is not changed during processing

const (
	HEAP_TYPE_MAX = iota
	HEAP_TYPE_MIN
)

type Heap struct {
	data     []int
	heapType int
}

func NewHeap(heapType int) (*Heap, error) {
	if heapType == HEAP_TYPE_MIN || heapType == HEAP_TYPE_MAX {
		return &Heap{[]int{}, heapType}, nil
	} else {
		return nil, errors.New("datastructure/Heap.newHeap(): Invalid heap type")
	}
}

func getParentIdx(index int) int { return (index - 1) / 2 }
func getLChildIdx(index int) int { return index*2 + 1 }
func getRChildIdx(index int) int { return index*2 + 2 }

func (h *Heap) getParentIdx(index int) int { return h.data[getParentIdx(index)] }
func (h *Heap) getLChildIdx(index int) int { return h.data[getLChildIdx(index)] }
func (h *Heap) getRChildIdx(index int) int { return h.data[getRChildIdx(index)] }

func (h *Heap) insert(data int) {
	h.data = append(h.data, data)
	h.heapify()
}

func (h *Heap) heapify() {
	nodeIdx := len(h.data) - 1
	if h.heapType == HEAP_TYPE_MAX {
		for h.data[nodeIdx] > h.data[getParentIdx(nodeIdx)] {
			h.swapData(nodeIdx, getParentIdx(nodeIdx))
			nodeIdx = getParentIdx(nodeIdx)
		}
	} else if h.heapType == HEAP_TYPE_MIN {
		for h.data[nodeIdx] < h.data[getParentIdx(nodeIdx)] {
			h.swapData(nodeIdx, getParentIdx(nodeIdx))
			nodeIdx = getParentIdx(nodeIdx)
		}
	}
}

func (h *Heap) extractTop() int {
	// Save top element to return it at the end: MAX or MIN depending on HASH_TYPE_*
	saveTop := h.data[0]

	// Swap last element with the top element that was popped
	lastIdx := len(h.data) - 1
	h.swapData(0, lastIdx)
	h.data = h.data[:lastIdx]

	// Reorder heap
	h.reorder(0)

	return saveTop
}

func (h *Heap) reorder(idx int) {
	lChild := getLChildIdx(idx)
	rChild := getRChildIdx(idx)

	if h.heapType == HEAP_TYPE_MAX {
		// no children
		if lChild > len(h.data)-1 {
			return
			// only left child
		} else if rChild > len(h.data)-1 && h.data[idx] < h.data[lChild] {
			h.swapData(idx, lChild)
		} else if h.data[lChild] > h.data[rChild] && h.data[idx] < h.data[lChild] {
			h.swapData(idx, lChild)
			h.reorder(lChild)
		} else if h.data[lChild] < h.data[rChild] && h.data[idx] < h.data[rChild] {
			h.swapData(idx, rChild)
			h.reorder(rChild)
		}
	} else if h.heapType == HEAP_TYPE_MIN {
		// no children
		if lChild > len(h.data)-1 {
			return
			// only left child
		} else if rChild > len(h.data)-1 && h.data[idx] > h.data[lChild] {
			h.swapData(idx, lChild)
		} else if h.data[lChild] < h.data[rChild] && h.data[idx] > h.data[lChild] {
			h.swapData(idx, lChild)
			h.reorder(lChild)
		} else if h.data[lChild] > h.data[rChild] && h.data[idx] > h.data[rChild] {
			h.swapData(idx, rChild)
			h.reorder(rChild)
		}
	}
}

func (h *Heap) swapData(idx1, idx2 int) {
	h.data[idx1], h.data[idx2] = h.data[idx2], h.data[idx1]
}
