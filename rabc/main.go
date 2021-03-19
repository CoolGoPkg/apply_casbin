package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("D:/mark/go/src/casbin_apply/rabc/model.conf", "D:/mark/go/src/casbin_apply/rabc/policy.csv")
	if err != nil {
		fmt.Print("err : ", err)
		return
	}

	sub := "berry" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.

	var ok bool
	ok, err = e.Enforce(sub, obj, act)

	if err != nil {
		fmt.Println("enforce err : ", err)
		// handle err
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("berry鉴权通过")
	} else {
		// deny the request, show an error
		fmt.Println("berry鉴权失败")
	}

	results, err := e.BatchEnforce([][]interface{}{{"berry", "data2", "write"},
		{"berry", "data1", "read"}, {"berry", "data2", "read"}})
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
