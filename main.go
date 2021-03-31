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
	obj := "data2" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.

	//增加susan为data2数据的admin角色
	added, err := e.AddGroupingPolicy("susan", "data2_admin")
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
