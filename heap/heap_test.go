package heap

import (
	"fmt"
	"testing"
)

func TestNewHeap(t *testing.T) {
	maxHeap, err := NewHeap(HEAP_TYPE_MAX)
	if err != nil && maxHeap.heapType != HEAP_TYPE_MAX {
		t.Errorf("- Unable to instantiate MAX Heap")
	}
	fmt.Println("+ MaxHeap creation ok:", maxHeap)

	minHeap, err := NewHeap(HEAP_TYPE_MIN)
	if err != nil && maxHeap.heapType != HEAP_TYPE_MIN {
		t.Errorf("- Unable to instantiate MIN Heap")
	}
	fmt.Println("+ MinHeap creation ok:", minHeap)

	wrongType, err := NewHeap(100)
	if err == nil && wrongType != nil {
		t.Errorf("- Should not allow any other heap type than HEAP_TYPE_MIN & HEAP_TYPE_MAX")
	}
	fmt.Println("+ Wrong heap type ok:", wrongType)
}

func TestInsertHeapMax(t *testing.T) {
	h, _ := NewHeap(HEAP_TYPE_MAX)
	h.insert(100)
	h.insert(200)
	h.insert(300)
	h.insert(10)
	h.insert(20)
	h.insert(30)
	h.insert(250)
	h.insert(350)

	expectedResult := []int{350, 300, 250, 100, 20, 30, 200, 10}
	for i := 0; i < len(h.data); i++ {
		if h.data[i] != expectedResult[i] {
			t.Errorf("- Incorrect result")
			fmt.Println("- Expected result:", expectedResult)
			fmt.Println("- Actual result:", h.data)
			return
		}
	}
	fmt.Println("+ MaxHeap insertion ok:", h.data)
}

func TestInsertHeapMin(t *testing.T) {
	h, _ := NewHeap(HEAP_TYPE_MIN)
	h.insert(100)
	h.insert(200)
	h.insert(300)
	h.insert(10)
	h.insert(20)
	h.insert(30)
	h.insert(250)
	h.insert(350)
	expectedResult := []int{10, 20, 30, 200, 100, 300, 250, 350}
	for i := 0; i < len(h.data); i++ {
		if h.data[i] != expectedResult[i] {
			t.Errorf("- Incorrect result")
			fmt.Println("- Expected result:", expectedResult)
			fmt.Println("- Actual result:", h.data)
			return
		}
	}
	fmt.Println("+ MinHeap insertion ok:", h.data)
}

func TestExtractTopHeapMax(t *testing.T) {
	h, _ := NewHeap(HEAP_TYPE_MAX)
	h.insert(100)
	h.insert(200)
	h.insert(300)
	h.insert(10)
	h.insert(20)
	h.insert(30)
	h.insert(250)
	h.insert(350)
	fmt.Println(h.data)

	topValue := h.extractTop()
	if topValue != 350 {
		t.Errorf("- Wrong MAX value extracted from the heap")
	}
	fmt.Println("+ Correct MAX value extracted from the heap:", topValue)

	expectedResult := []int{300, 100, 250, 10, 20, 30, 200}
	for i := 0; i < len(h.data); i++ {
		if h.data[i] != expectedResult[i] {
			t.Errorf("- Incorrect result")
			fmt.Println("- Expected result:", expectedResult)
			fmt.Println("- Actual result:", h.data)
			return
		}
	}
	fmt.Println("+ Correct resulting heap after extracting top value:", h.data)
}

func TestExtractTopHeapMin(t *testing.T) {
	h, _ := NewHeap(HEAP_TYPE_MIN)
	h.insert(100)
	h.insert(200)
	h.insert(300)
	h.insert(10)
	h.insert(20)
	h.insert(30)
	h.insert(250)
	h.insert(350)
	fmt.Println(h.data)

	topValue := h.extractTop()
	if topValue != 10 {
		t.Errorf("- Wrong MIN value extracted from the heap")
	}
	fmt.Println("+ Correct MIN value extracted from the heap:", topValue)

	expectedResult := []int{20, 100, 30, 200, 350, 300, 250}
	for i := 0; i < len(h.data); i++ {
		if h.data[i] != expectedResult[i] {
			t.Errorf("- Incorrect result")
			fmt.Println("- Expected result:", expectedResult)
			fmt.Println("- Actual result:", h.data)
			return
		}
	}
}
