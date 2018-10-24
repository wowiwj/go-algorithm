package shell

import "github.com/wowiwj/go-algorithm/sort"

func Sort(a sort.Sortable)  {
	N := a.Len()
	h := 1
	for h < N / 3 {
		h = 3 * h + 1
	}

	for h > 1  {
		for i := h; i < N;i++ {
			for j := i;j >= h && a.Less(j,j-h);j-=h {
				a.Swap(j,j-h)
			}
		}
		h = h / 3
	}
}
