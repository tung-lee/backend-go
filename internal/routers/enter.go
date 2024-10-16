package routers

import (
	"backend-go/internal/routers/admin"
	"backend-go/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)