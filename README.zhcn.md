[For English](/README.md)
# SqlExecutor-Go
这是一个动态查询客户端的Golang实现。你可以利用此组件构建动态查询参数。

## 持续集成 
![](https://travis-ci.org/FeiniuBus/SqlExecutor-Go.svg?branch=master)

## 示例
* 在可以在 `./samples` 中找到示例代码

***使用前需要导入程序包***
```go
import (
	DynamicQuery "github.com/FeiniuBus/SqlExecutor-Go/dynamicquery"
)
```
***然后就可以使用了***
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

## 安装方式
`go get -u github.com/FeiniuBus/SqlExecutor-Go/dynamicquery`
