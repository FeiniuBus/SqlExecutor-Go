package dynamicquery

// QueryOperation ...
type QueryOperation int

const (
	_ QueryOperation = iota
	//Equal 等于
	Equal = 1
	// LessThan 小于
	LessThan = 2
	// LessThanOrEqual 小于等于
	LessThanOrEqual = 3
	// GreaterThan 大于
	GreaterThan = 4
	// GreaterThanOrEqual 大于等于
	GreaterThanOrEqual = 5
	// Contains 包含(like)
	Contains = 6
	// StartsWith 以xx开始
	StartsWith = 7
	// EndsWith 以xx结束
	EndsWith = 8
	// In 包含在xx里
	In = 9
	// Any ...
	Any = 10
	// DateTimeLessThanOrEqualThenDay ...
	DateTimeLessThanOrEqualThenDay = 11
)
