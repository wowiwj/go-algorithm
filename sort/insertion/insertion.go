package insertion

import "github.com/wowiwj/go-algorithm/sort"



func Sort(a sort.Sortable)  {
	N := a.Len()

	for i := 1; i < N ; i++  {
		for j := i;j > 0 && a.Less(j,j-1);j-- {
			a.Swap(j,j-1)
		}
	}
}
