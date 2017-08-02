package main

/*
Find the integer part of the minimum average waiting time.
See https://www.hackerrank.com/challenges/minimum-average-waiting-time/problem
*/

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	avg := MinAvgTime(os.Stdin)
	fmt.Println(avg)
}

func MinAvgTime(r io.Reader) uint64 {
	s := bufio.NewScanner(r)

	var n int
	if s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)
		n, _ = strconv.Atoi(line)
	}

	q := make([]*Customer, 0, n)

	for s.Scan() {
		line := s.Text()
		c := readCustomer(line)
		q = append(q, c)
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].Time < q[j].Time
	})

	pq := make(PriorityQueue, 0, n)
	heap.Init(&pq)

	var (
		sum        uint64
		totalCookT uint64
	)
	for {
		for i := 0; len(q) > 0; i++ {
			c := q[0]
			if c.Time > totalCookT && i > 0 {
				break
			}
			heap.Push(&pq, c)
			q = q[1:]
			//fmt.Println("next to queue", nextT, t)
		}

		if pq.Len() == 0 {
			break
		}
		next := heap.Pop(&pq).(*Customer)
		if totalCookT == 0 {
			totalCookT = next.Time
		}

		extraT := uint64(0)
		if totalCookT > next.Time {
			extraT = totalCookT - next.Time
		}
		totalCookT += next.CookDuration
		sum += next.CookDuration + extraT

		//fmt.Println("next", next.Time, next.CookDuration, extraT, totalCookT, "sum", sum)
	}

	return sum / uint64(n)
}

func readCustomer(line string) *Customer {
	line = strings.TrimSpace(line)
	parts := strings.SplitN(line, " ", 2)
	time, _ := strconv.ParseUint(parts[0], 10, 64)
	duration, _ := strconv.ParseUint(parts[1], 10, 64)

	return &Customer{time, duration}
}

type Customer struct {
	Time         uint64
	CookDuration uint64
}

func NewCustomer(t, l string) *Customer {
	time, _ := strconv.ParseUint(t, 10, 64)
	duration, _ := strconv.ParseUint(l, 10, 64)

	return &Customer{time, duration}
}

type PriorityQueue []*Customer

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].CookDuration < pq[j].CookDuration
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	customer := x.(*Customer)
	*pq = append(*pq, customer)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
