package search

import (
	"fmt"
	"sort"
)

func BinarySearchLookup(key float64, st *SortedTable) (offsets []int, err error) {
	if !sort.SliceIsSorted(st.Keys, func(i, j int) bool { return st.Keys[i] < st.Keys[j] }) {
		panic("not sorted!!")
	}
	i := sort.SearchFloat64s(st.Keys, key)
	for ; i < len(st.Keys); i++ {
		if st.Keys[i] > key {
			break
		} else if st.Keys[i] == key {

			offsets = append(offsets, st.Offsets[i])
		}
	}

	if len(offsets) > 0 {
		return offsets, nil
	}
	return nil, fmt.Errorf("The following key <%f> is not found in the index", key)
}
