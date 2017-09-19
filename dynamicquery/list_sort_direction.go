package dynamicquery

// ListSortDirection ...
type ListSortDirection int

const (
	_ ListSortDirection = iota
	// Ascending 正序
	Ascending = 0
	// Descending 倒序
	Descending = 1
)
