package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, _ := gormadapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/casbin", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("./model.conf", a)
	// e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	// Load the policy from DB.
	e.LoadPolicy()

	sub := "susan" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	//add susan
	added, err := e.AddPolicy("susan", "data2", "read")
	fmt.Println(added)
	fmt.Println(err)

	//find susan
	filteredPolicy := e.GetFilteredPolicy(0, "susan")
	fmt.Println(filteredPolicy)

	//update susan
	updated, err := e.UpdatePolicy([]string{"susan", "data2", "read"}, []string{"susan", "data2", "write"})
	fmt.Println(updated)

	//delete susan
	removed, err := e.RemovePolicy("susan", "data2", "write")
	fmt.Println("remove susan:", removed)

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
