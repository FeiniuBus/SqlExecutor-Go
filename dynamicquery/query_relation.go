package dynamicquery

//QueryRelation ...
type QueryRelation int

const (
	_ QueryRelation = iota
	//And Relation
	And = 0
	//Or Relation
	Or = 1
)
