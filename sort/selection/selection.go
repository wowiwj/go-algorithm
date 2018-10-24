package selection

import "github.com/wowiwj/go-algorithm/sort"

func Sort(a sort.Sortable)  {
	N := a.Len()
	for i := 0;i < N;i++  {
		min := i;
		for j := i + 1; j < N; j++ {
			if a.Less(j,min) {
				min = j
			}
		}
		a.Swap(i,min)
	}
}
