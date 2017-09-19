package dynamicquery

// Param 动态查询参数.
type Param struct {
	Field    string
	Operator QueryOperation
	Value    interface{}
}

// NewParam instant a new Param
func NewParam(field string, operator QueryOperation, value interface{}) *Param {
	return &Param{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}
