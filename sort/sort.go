package sort

type Sortable interface {
	// len is number of elements in collection
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i,j int) bool
	// swap
	Swap(i,j int)
}

func IsSorted(a Sortable) bool  {
	for i := 1; i < a.Len() ; i++  {
		if a.Less(i,i-1) {
			return false
		}
	}
	return true
}
