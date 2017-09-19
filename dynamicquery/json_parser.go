package dynamicquery

// JSONParser ...
type JSONParser struct {
}

// NewJSONParser ...
func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

// ParseQueryOperation ...
func (parser *JSONParser) ParseQueryOperation(operator QueryOperation) string {
	switch operator {
	case Equal:
		return "Equal"
	case LessThan:
		return "LessThan"
	case LessThanOrEqual:
		return "LessThanOrEqual"
	case GreaterThan:
		return "GreaterThan"
	case GreaterThanOrEqual:
		return "GreaterThanOrEqual"
	case Contains:
		return "Contains"
	case StartsWith:
		return "StartsWith"
	case EndsWith:
		return "EndsWith"
	case In:
		return "In"
	case Any:
		return "Any"
	case DateTimeLessThanOrEqualThenDay:
		return "DateTimeLessThanOrEqualThenDay"
	}
	return "Equal"
}

// ParseListSortDirection ...
func (parser *JSONParser) ParseListSortDirection(val ListSortDirection) string {
	switch val {
	case Ascending:
		return "Ascending"
	case Descending:
		return "Descending"
	}
	return "Ascending"
}

// ParseQueryRelation ...
func (parser *JSONParser) ParseQueryRelation(val QueryRelation) string {
	switch val {
	case And:
		return "And"
	case Or:
		return "Or"
	}
	return "And"
}
