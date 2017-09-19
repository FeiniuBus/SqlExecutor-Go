package dynamicquery

// ParamGroup ...
type ParamGroup struct {
	Relation    QueryRelation
	ChildGroups []*ParamGroup
	Params      []*Param
}

// NewParamGroup ...
func NewParamGroup(relation QueryRelation) *ParamGroup {
	return &ParamGroup{
		Relation:    relation,
		ChildGroups: make([]*ParamGroup, 0),
		Params:      make([]*Param, 0),
	}
}

// CheckDynamicQueryParamGroup ...
func (group *ParamGroup) CheckDynamicQueryParamGroup() error {
	if len(group.Params) > 0 && len(group.ChildGroups) > 0 {
		return NewException("DynamicQueryParamGroup不能同时添加Params和Group")
	}
	return nil
}
