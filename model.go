package main

import "sort"

// ListItem interface
type ListItem interface {
	Hero | Item
	GetID() int
}

// SortList func
func SortList[I ListItem](l []I) {
	sort.Slice(l, func(i, j int) bool {
		return l[i].GetID() < l[j].GetID()
	})
}
