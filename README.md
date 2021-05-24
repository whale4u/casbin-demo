# casbin-demo
- feature/v1 基本实现casbin
- feature/v2 实现基于gorm+mysql的casbin增删查改
- feature/v3 相较于v2增加model角色概念
- feature/v4 相较于v3增加自定义函数
- feature/v5 实现RBAC版本casbin

## V5 RBAC Casbin
```
go run main.go
alice /data2_admin/xxx GET  Pass
```