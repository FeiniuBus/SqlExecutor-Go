[中文文档](/README.zhcn.md)
# SqlExecutor-Go
Dynamic Query Client coden by golang. You could build a Dynamic Query Parameter very simply with it.

## Pipeline 
[![](https://travis-ci.org/FeiniuBus/SqlExecutor-Go.svg?branch=master)](https://travis-ci.org/FeiniuBus/SqlExecutor-Go)

## Dynamic Query Sample
* You can find this code file in `./samples`

***IMPORT PACKAGE BEFORE USE***
```go
import (
	DynamicQuery "github.com/FeiniuBus/SqlExecutor-Go/dynamicquery"
)
```
***Then let's rock***
```go
builder := DynamicQuery.NewBuilder(true)
child1, err := builder.ParamGroupBuilder.CreateChildAndGroup()
if err != nil {
	return nil, err
}
child1.ParamBuilder.Any("Extra",
	func(sub *DynamicQuery.ParamGroupBuilder) {
		sub.ParamBuilder.Equal("Guest", "Andy")
	})
child2, err := builder.ParamGroupBuilder.CreateChildOrGroup()
if err != nil {
	return nil, err
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
builder.OrderBy("Amout", DynamicQuery.Ascending).Select("Guest").Take(10).Skip(10)
dynamicQuery := builder.Build()
```

## Package install
`go get -u github.com/FeiniuBus/SqlExecutor-Go/dynamicquery`
