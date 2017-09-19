package dynamicquery

// DynamicQuery ...
type DynamicQuery struct {
	ParamGroup *ParamGroup
	Order      []*Order
	Skip       int
	Take       int
	Pager      bool
	Select     string
}

// NewDynamicQuery ...
func NewDynamicQuery(pager bool) *DynamicQuery {
	return &DynamicQuery{
		ParamGroup: NewParamGroup(And),
		Order:      make([]*Order, 0),
		Skip:       0,
		Take:       10,
		Pager:      pager,
		Select:     "",
	}
}
