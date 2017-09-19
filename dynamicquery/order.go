package dynamicquery

// Order 排序设置
type Order struct {
	Name string
	Sort ListSortDirection
}

// NewOrder ...
func NewOrder(name string, sort ListSortDirection) *Order {
	return &Order{
		Name: name,
		Sort: sort,
	}
}
