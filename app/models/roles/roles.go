package roles

import (
	"gin-admin/app/utility/app"
)

func GetRole() (entity []Entity, err error) {
	err = app.DB().Find(&entity)
	return
}

func GetRoleById(id int) (entity Entity, err error) {
	_, err = app.DB().Where("id=?", id).Get(&entity)
	return
}
