package selection

import (
	"fmt"
	"github.com/wowiwj/go-algorithm/sort"
	"testing"
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

func TestSort(t *testing.T) {
	
	arr := intSlice{1,2,1,2,5,3,2,1}
	Sort(arr)
	for i := 0; i < len(arr);i++{
		fmt.Println(arr[i])
	}

	if ! sort.IsSorted(arr) {
		t.Error("fail")
	}
}