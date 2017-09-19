package dynamicquery

// JSONParam ...
type JSONParam struct {
	Field    string      `json:"Field"`
	Operator string      `json:"Operator"`
	Value    interface{} `json:"Value"`
}

// NewJSONParam ...
func NewJSONParam(param *Param) *JSONParam {
	parser := NewJSONParser()
	return &JSONParam{
		Field:    param.Field,
		Operator: parser.ParseQueryOperation(param.Operator),
		Value:    param.Value,
	}
}
