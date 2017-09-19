package dynamicquery

// JSONOrder ...
type JSONOrder struct {
	Name string `json:"Name"`
	Sort string `json:"Sort"`
}

// NewJSONOrder ...
func NewJSONOrder(order *Order) *JSONOrder {
	parser := NewJSONParser()
	return &JSONOrder{
		Name: order.Name,
		Sort: parser.ParseListSortDirection(order.Sort),
	}
}
