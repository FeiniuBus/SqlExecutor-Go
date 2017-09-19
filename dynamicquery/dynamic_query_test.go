package dynamicquery

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDynamicQuery(t *testing.T) {
	builder := NewBuilder(true)

	child1, err := builder.ParamGroupBuilder.CreateChildAndGroup()
	if err != nil {
		t.Error(err)
	}
	child1.ParamBuilder.Any("Extra",
		func(sub *ParamGroupBuilder) {
			sub.ParamBuilder.Equal("Guest", "Andy")
		})
	child2, err := builder.ParamGroupBuilder.CreateChildOrGroup()
	if err != nil {
		t.Error(err)
	}
	child2.ParamBuilder.Contains("Address", "chengdu")
	child2.ParamBuilder.EndsWith("Address", "lnk")
	child2.ParamBuilder.Equal("Disabled", false)
	child2.ParamBuilder.GreaterThan("Amout", 10)
	child2.ParamBuilder.GreaterThanOrEqual("Price", 100)
	child2.ParamBuilder.In("Drink", "mileshake,coffee")
	child2.ParamBuilder.LessThan("Count", 10)
	child2.ParamBuilder.LessThanOrEqual("Total", 100)
	child2.ParamBuilder.StartsWith("Url", "Http://")

	builder.OrderBy("Amout", Ascending).Select("Guest").Take(10).Skip(10)

	query := builder.Build()
	j, err := json.Marshal(query)
	output := string(j)
	fmt.Print(output)
}
