package controller

import (
	"github.com/taoshihan1991/imaptool/models"
)

func CheckKefuPass(username string, password string) (models.User, models.User_role, bool) {
	info := models.FindUser(username)
	var uRole models.User_role
	uRole = models.FindRoleByUserId(info.ID)

	return info, uRole, true
}
