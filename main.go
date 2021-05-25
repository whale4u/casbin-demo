package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	// 新建1个文件适配器
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")

	sub := "susan" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok1, err := e.Enforce(sub, obj, act)
	if err != nil {
		// handle err
		fmt.Printf("%s", err)
	}

	if ok1 {
		// permit alice to read data1
		fmt.Println("check pass")
	} else {
		// deny the request, show an error
		fmt.Println("check fail")
	}

	// 新增1条policy
	added, err := e.AddPolicy("susan", "data1", "read")
	fmt.Println("added operation: ", added)
	// fmt.Println(err)

	ok2, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
		fmt.Printf("%s", err)
	}

	if ok2 {
		// permit alice to read data1
		fmt.Println("check pass")
	} else {
		// deny the request, show an error
		fmt.Println("check fail")
	}
}
