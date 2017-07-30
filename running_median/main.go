package main

/*
Given an input stream of N integers, you must perform the following task for each i-th integer:

1. Add the ith integer to a running list of integers.
2. Find the median of the updated list (i.e., for the first element through the ith element).
3. Print the list's updated median on a new line. The printed value must be a double-precision number
scaled to decimal place (i.e., 12.3 format).

See https://www.hackerrank.com/challenges/find-the-running-median/problem
*/

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	var n int
	if s.Scan() {
		sn := s.Text()
		sn = strings.TrimSpace(sn)
		n, _ = strconv.Atoi(sn)
		_ = n
	}

	left := NewMaxHeap(nil)
	right := NewMinHeap(nil)

	var (
		i int
		m float64
	)
	for s.Scan() {
		sn := s.Text()
		sn = strings.TrimSpace(sn)
		v, _ := strconv.Atoi(sn)

		if i == 0 {
			// the first median is always a0
			m = float64(v)
		} else if i%2 != 0 {
			// for even number of elements, the median is the average of the two middle elements of the sorted sample
			if float64(v) > m {
				heap.Push(left, int(m))
				heap.Push(right, v)
			} else {
				heap.Push(left, v)
				heap.Push(right, int(m))
			}
			m = float64((left.IntHeap[0] + right.IntHeap[0])) / 2.
		} else {
			// for odd number of elements, the median is the middle element of the sorted sample
			if float64(v) > m {
				heap.Push(right, v)
				m = float64(heap.Pop(right).(int))
			} else if float64(v) < m {
				heap.Push(left, v)
				m = float64(heap.Pop(left).(int))
			} else {
				m = float64(v)
			}
		}
		i++
		fmt.Printf("%.1f\n", m)
	}
}

type IntHeap []int

func (h IntHeap) Len() int      { return len(h) }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Heap struct {
	IntHeap
	less func([]int, int, int) bool
}

func (h Heap) Less(i, j int) bool { return h.less([]int(h.IntHeap), i, j) }

func NewMinHeap(v []int) *Heap {
	return &Heap{
		IntHeap: IntHeap(v),
		less:    func(h []int, i, j int) bool { return h[i] < h[j] },
	}
}

func NewMaxHeap(v []int) *Heap {
	return &Heap{
		IntHeap: IntHeap(v),
		less:    func(h []int, i, j int) bool { return h[i] > h[j] },
	}
}
