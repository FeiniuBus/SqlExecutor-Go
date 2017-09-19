package dynamicquery

// JSONDynamicQuery ...
type JSONDynamicQuery struct {
	ParamGroup *JSONParamGroup `json:"ParamGroup"`
	Order      []*JSONOrder    `json:"Order"`
	Skip       int             `json:"Skip"`
	Take       int             `json:"Take"`
	Pager      bool            `json:"Pager"`
	Select     string          `json:"Select"`
}

// NewJSONDynamicQuery ...
func NewJSONDynamicQuery(dynamicQuery *DynamicQuery) *JSONDynamicQuery {
	returnValue := &JSONDynamicQuery{
		ParamGroup: NewJSONParamGroup(dynamicQuery.ParamGroup),
		Skip:       dynamicQuery.Skip,
		Take:       dynamicQuery.Take,
		Pager:      dynamicQuery.Pager,
		Select:     dynamicQuery.Select,
		Order:      make([]*JSONOrder, 0),
	}
	for _, item := range dynamicQuery.Order {
		returnValue.Order = append(returnValue.Order, NewJSONOrder(item))
	}
	return returnValue
}
