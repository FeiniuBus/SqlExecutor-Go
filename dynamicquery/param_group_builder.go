package dynamicquery

// ParamGroupBuilder ...
type ParamGroupBuilder struct {
	paramGroup   *ParamGroup
	ParamBuilder *ParamBuilder
}

// NewParamGroupBuilder ...
func NewParamGroupBuilder() *ParamGroupBuilder {
	paramGroup := NewParamGroup(And)
	return &ParamGroupBuilder{
		paramGroup:   paramGroup,
		ParamBuilder: NewParamBuilderWithParamGroup(paramGroup),
	}
}

// NewParamGroupBuilderWithParamGroup ...
func NewParamGroupBuilderWithParamGroup(paramGroup *ParamGroup) *ParamGroupBuilder {
	return &ParamGroupBuilder{
		paramGroup:   paramGroup,
		ParamBuilder: NewParamBuilderWithParamGroup(paramGroup),
	}
}

// Relation ...
func (builder *ParamGroupBuilder) Relation(relation QueryRelation) *ParamGroupBuilder {
	builder.paramGroup.Relation = relation
	return builder
}

// RelationAnd ...
func (builder *ParamGroupBuilder) RelationAnd() *ParamGroupBuilder {
	return builder.Relation(And)
}

// RelationOr ..
func (builder *ParamGroupBuilder) RelationOr() *ParamGroupBuilder {
	return builder.Relation(Or)
}

// CreateChildGroup ...
func (builder *ParamGroupBuilder) CreateChildGroup(relation QueryRelation) (*ParamGroupBuilder, error) {
	err := builder.paramGroup.CheckDynamicQueryParamGroup()
	if err != nil {
		return nil, err
	}
	childGroup := NewParamGroup(relation)
	builder.paramGroup.ChildGroups = append(builder.paramGroup.ChildGroups, childGroup)
	return NewParamGroupBuilderWithParamGroup(childGroup), nil
}

// CreateChildAndGroup ...
func (builder *ParamGroupBuilder) CreateChildAndGroup() (*ParamGroupBuilder, error) {
	return builder.CreateChildGroup(And)
}

// CreateChildOrGroup ...
func (builder *ParamGroupBuilder) CreateChildOrGroup() (*ParamGroupBuilder, error) {
	return builder.CreateChildGroup(Or)
}

// Build ...
func (builder *ParamGroupBuilder) Build() *ParamGroup {
	return builder.paramGroup
}
