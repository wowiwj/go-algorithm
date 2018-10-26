package sort

import (
	"math/rand"
	"testing"
	"time"
)
type intSlice []int

func (arr intSlice) Len() int {
	return len(arr)
}

func (arr intSlice) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr intSlice) Swap(i, j int) {
	arr[i],arr[j] = arr[j],arr[i]
}


func runTime(sortFunc func(Sortable),a Sortable) time.Duration  {
	t1 := time.Now()
	sortFunc(a)
	t2 := time.Now()
	return t2.Sub(t1)
}

func timeRandomInput(sortFunc func(Sortable),N,T int) time.Duration  {
	total := time.Microsecond * 0
	a := intSlice{}
	for t := 0; t < T ; t++  {
		for i := 0; i < N;i++ {
			a = append(a,rand.Int())
		}
		total += runTime(sortFunc,a)
	}
	return total
}

func Test(t *testing.T) {
	
}