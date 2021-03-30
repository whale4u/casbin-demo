package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	sub := "susan" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	//add susan
	added, err := e.AddPolicy("susan", "data1", "read")
	fmt.Println(added)
	fmt.Println(err)

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
