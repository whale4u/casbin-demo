package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, err := gormadapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/casbin", true) // Your driver and data source.
	fmt.Println(err)
	e, _ := casbin.NewEnforcer("./model.conf", a)
	// e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	// 注册自定义函数
	e.AddFunction("my_func", KeyMatchFunc)

	// Load the policy from DB.
	e.LoadPolicy()

	sub := "susan" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	_, err1 := e.AddPolicy("data2_admin", "data2", "read")
	fmt.Println(err1)
	_, err2 := e.AddPolicy("data2_admin", "data2", "write")
	fmt.Println(err2)
	//增加susan为data2数据的admin角色
	added, err := e.AddGroupingPolicy("susan", "data2_admin")
	fmt.Println(added)
	fmt.Println(err)

	ok, err := e.Enforce(sub, obj, act)

	// fmt.Println(ok)

	if err != nil {
		// handle err
		fmt.Printf("%s", err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println(sub, obj, act, " Pass")
	} else {
		// deny the request, show an error
		fmt.Println(sub, obj, act, " Fail")
	}
}

func KeyMatch(key1 string, key2 string) bool {
	return key1 == key2
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
