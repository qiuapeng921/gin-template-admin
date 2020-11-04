package database

import (
	"fmt"
	"gin-admin/app/models/admin_role"
	"gin-admin/app/models/admins"
	"gin-admin/app/models/menus"
	"gin-admin/app/models/permissions"
	"gin-admin/app/models/role_permission"
	"gin-admin/app/models/roles"
	"gin-admin/app/models/system_config"
	"gin-admin/app/utility/app"
	"gin-admin/app/utility/system"
	"xorm.io/xorm"
)

var orm *xorm.Engine

func AutoGenTable() {
	orm = app.DB()
	_ = orm.Sync2(
		&admins.Entity{},
		&admin_role.Entity{},
		&menus.Entity{},
		&permissions.Entity{},
		&roles.Entity{},
		&role_permission.Entity{},
		&system_config.Entity{},
	)

	if result, _ := orm.IsTableEmpty(&admins.Entity{}); result {
		defaultData()
	}
}

func defaultData() {
	var admin admins.Entity
	admin.Username = "admin"
	admin.Password = system.EncodeMD5("123456")
	admin.Phone = "15249279779"
	if _, err := orm.InsertOne(&admin); err != nil {
		fmt.Println("初始化超管失败," + err.Error())
	}
	fmt.Println("初始化超管成功")

	var role roles.Entity
	role.RoleName = "超管"
	role.RoleDesc = "超级管理员"
	if _, err := orm.Insert(&role); err != nil {
		fmt.Println("初始化角色失败," + err.Error())
	}
	fmt.Println("初始化角色成功")

	var adminRole admin_role.Entity
	adminRole.AdminId = 1
	adminRole.RoleId = 1
	if _, err := orm.Insert(&adminRole); err != nil {
		fmt.Println("初始化用户角色关系失败," + err.Error())
	}
	fmt.Println("初始化用户角色关系成功")
}
