package dynamicquery

// JSONParamGroup ...
type JSONParamGroup struct {
	Relation    string            `json:"Relation"`
	ChildGroups []*JSONParamGroup `json:"ChildGroup"`
	Params      []*JSONParam      `json:"Params"`
}

// NewJSONParamGroup ...
func NewJSONParamGroup(group *ParamGroup) *JSONParamGroup {
	parser := NewJSONParser()
	returnValue := &JSONParamGroup{
		Relation:    parser.ParseQueryRelation(group.Relation),
		ChildGroups: make([]*JSONParamGroup, 0),
		Params:      make([]*JSONParam, 0),
	}
	for _, item := range group.ChildGroups {
		returnValue.ChildGroups = append(returnValue.ChildGroups, NewJSONParamGroup(item))
	}
	for _, item := range group.Params {
		returnValue.Params = append(returnValue.Params, NewJSONParam(item))
	}
	return returnValue
}
