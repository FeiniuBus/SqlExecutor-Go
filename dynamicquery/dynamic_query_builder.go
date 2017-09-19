package dynamicquery

// Builder ...
type Builder struct {
	query             *DynamicQuery
	ParamGroupBuilder *ParamGroupBuilder
}

// NewBuilder ...
func NewBuilder(pager bool) *Builder {
	query := NewDynamicQuery(pager)
	return &Builder{
		query:             query,
		ParamGroupBuilder: NewParamGroupBuilderWithParamGroup(query.ParamGroup),
	}
}

// Skip ...
func (builder *Builder) Skip(skip int) *Builder {
	builder.query.Skip = skip
	return builder
}

// Take ...
func (builder *Builder) Take(take int) *Builder {
	builder.query.Take = take
	return builder
}

// Pager ...
func (builder *Builder) Pager(pager bool) *Builder {
	builder.query.Pager = pager
	return builder
}

// Select ...
func (builder *Builder) Select(selectStr string) *Builder {
	builder.query.Select = selectStr
	return builder
}

// OrderBy ...
func (builder *Builder) OrderBy(name string, sort ListSortDirection) *Builder {
	hasElement := false
	for _, item := range builder.query.Order {
		if item.Name == name && item.Sort == sort {
			hasElement = true
			break
		}
	}
	if hasElement == false {
		builder.query.Order = append(builder.query.Order, NewOrder(name, sort))
	}
	return builder
}

// Build ...
func (builder *Builder) Build() *JSONDynamicQuery {
	return NewJSONDynamicQuery(builder.query)
}
