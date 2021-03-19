package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

// 	属性鉴权
type subject struct {
	Owner string
	Age   int
	Name  string
	Title string
}

//  ABAC是 基于属性的访问控制，可以使用主体、客体或动作的属性，
// 而不是字符串本身来控制访问。 您之前可能就已经听过 XACML ，
// 是一个复杂的 ABAC 访问控制语言。 与XACML相比，
//Casbin的ABAC非常简单: 在ABAC中，
//可以使用struct(或基于编程语言的类实例) 而不是字符串来表示模型元素。
func main() {
	e, err := casbin.NewEnforcer("D:/mark/go/src/casbin_apply/abac/model.conf", "D:/mark/go/src/casbin_apply/abac/policy.csv")
	if err != nil {
		fmt.Print("err : ", err)
		return
	}

	obj := "/data3" // the user that wants to access a resource.
	sub := subject{
		Owner: "mark",
		Title: "admin",
		Age:   13,
	} // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.

	var ok bool
	ok, err = e.Enforce(sub, obj, act)

	if err != nil {
		fmt.Println("enforce err : ", err)
		// handle err
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("鉴权通过")
	} else {
		// deny the request, show an error
		fmt.Println("鉴权失败")
	}

	results, err := e.BatchEnforce([][]interface{}{{subject{
		Age: 19,
	}, "/data1", "read"}, {subject{
		Age: 100,
	}, "/data2", "write"}, {subject{
		Age: 18,
	}, "/data2", "read"}})
	if err != nil {
		fmt.Println("err : results err", err)
		return
	}
	for i, r := range results {
		var desc string
		if r {
			desc = "鉴权通过"

		} else {

			desc = "鉴权失败"
		}
		fmt.Printf("第%d组: %s \n", i, desc)
	}
}
