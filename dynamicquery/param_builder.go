package dynamicquery

import (
	"encoding/json"
	"time"
)

// ParamBuilder ...
type ParamBuilder struct {
	params []*Param
	group  *ParamGroup
}

// NewParamBuilder ...
func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{
		params: make([]*Param, 0),
	}
}

// NewParamBuilderWithParamGroup ...
func NewParamBuilderWithParamGroup(group *ParamGroup) *ParamBuilder {
	return &ParamBuilder{
		group: group,
	}
}

// Add ...
func (builder *ParamBuilder) Add(param *Param) *ParamBuilder {
	if builder.group != nil {
		builder.group.Params = append(builder.group.Params, param)
	} else {
		builder.params = append(builder.params, param)
	}
	return builder
}

// CreateParam ...
func (builder *ParamBuilder) CreateParam(opt QueryOperation, field string, value interface{}) *Param {
	return &Param{
		Operator: opt,
		Field:    field,
		Value:    value,
	}
}

// Equal ...
func (builder *ParamBuilder) Equal(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(Equal, field, value))
}

// LessThan ...
func (builder *ParamBuilder) LessThan(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(LessThan, field, value))
}

// LessThanOrEqual ...
func (builder *ParamBuilder) LessThanOrEqual(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(LessThanOrEqual, field, value))
}

// GreaterThan ...
func (builder *ParamBuilder) GreaterThan(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(GreaterThan, field, value))
}

// GreaterThanOrEqual ...
func (builder *ParamBuilder) GreaterThanOrEqual(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(GreaterThanOrEqual, field, value))
}

// Contains ...
func (builder *ParamBuilder) Contains(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(Contains, field, value))
}

// StartsWith ...
func (builder *ParamBuilder) StartsWith(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(StartsWith, field, value))
}

// EndsWith ...
func (builder *ParamBuilder) EndsWith(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(EndsWith, field, value))
}

// In ...
func (builder *ParamBuilder) In(field string, value interface{}) *ParamBuilder {
	return builder.Add(builder.CreateParam(In, field, value))
}

// Any ...
func (builder *ParamBuilder) Any(field string, predicate func(*ParamGroupBuilder)) *ParamBuilder {
	group := NewParamGroup(And)
	bu := NewParamGroupBuilderWithParamGroup(group)
	predicate(bu)
	value, _ := json.Marshal(group)
	return builder.Add(builder.CreateParam(Any, field, string(value)))
}

// DataTimeLessThanOrEqualThenDay ...
func (builder *ParamBuilder) DataTimeLessThanOrEqualThenDay(field string, value time.Time) *ParamBuilder {
	return builder.Add(builder.CreateParam(DateTimeLessThanOrEqualThenDay, field, value))
}
