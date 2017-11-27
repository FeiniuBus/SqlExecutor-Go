package dynamicquery

//QueryRelation ...
type QueryRelation int

const (
	_ QueryRelation = iota
	//And Relation
	And
	//Or Relation
	Or
	//AndNot Relation
	AndNot
	//OrNot relation
	OrNot
)
