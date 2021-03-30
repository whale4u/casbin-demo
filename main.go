package main

import (
	"fmt"
)

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	sub := "zhangsan" // the user that wants to access a resource.
	obj := "data1"    // the resource that is going to be accessed.
	act := "read"     // the operation that the user performs on the resource.

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
		fmt.Printf("%s", err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("Pass")
	} else {
		// deny the request, show an error
		fmt.Println("Fail")
	}
}
