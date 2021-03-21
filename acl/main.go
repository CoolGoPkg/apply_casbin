package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("D:/mark/go/src/casbin_apply/acl/model.conf",
		"D:/mark/go/src/casbin_apply/acl/policy.csv")
	if err != nil {
		fmt.Print("err : ", err)
		return
	}
	var ok bool
	//ok, err := e.AddPolicies([][]string{
	//	{"bob", "data1", "write"},
	//})
	//if err != nil {
	//	fmt.Print("add policy err : ", err)
	//	return
	//}
	//
	//if !ok {
	//	fmt.Print("add  : ", ok)
	//	return
	//}

	sub := "zhankun" // the user that wants to access a resource.
	obj := "source1" // the resource that is going to be accessed.
	act := "read"    // the operation that the user performs on the resource.

	ok, err = e.Enforce(sub, obj, act)

	if err != nil {
		fmt.Println("enforce err : ", err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("zhankun鉴权通过")
	} else {
		// deny the request, show an error
		fmt.Println("zhankun鉴权失败")
	}

	// You could use BatchEnforce() to enforce some requests in batches.
	// This method returns a bool slice, and this slice's index corresponds to the row index of the two-dimensional array.
	// e.g. results[0] is the result of {"alice", "data1", "read"}
	results, err := e.BatchEnforce([][]interface{}{
		{"alice", "data1", "read"},
		{"bob", "data1", "write"},
		{"jack", "data3", "read"}})
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
