package routers

import (
	v1 "admin-api/internal/api/v1"
	"admin-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)
	r.Use(middleware.Access)

	r.POST("login", v1.Auth.Login)
	r.POST("renewToken", v1.Auth.RenewToken)

	api := r.Group("")
	api.Use(middleware.Auth)
	{
		authApi := api.Group("auth")
		{
			authApi.GET("jwt", v1.Auth.JwtInfo)
			authApi.GET("menu", v1.Auth.MenuTree)
			authApi.POST("password", v1.Auth.ChangePwd)
		}

		groupApi := api.Group("group")
		{
			groupApi.GET("", v1.Auth.Group)
			groupApi.POST("", v1.Auth.AddGroup)
			groupApi.POST("edit", v1.Auth.ChangeGroup)
			groupApi.POST("user", v1.Auth.SetGroupUser)
			groupApi.POST("menu", v1.Auth.SetGroupMenu)
		}

		menuApi := api.Group("menu")
		{
			menuApi.GET("", v1.Auth.Menu)
			menuApi.GET("info", v1.Auth.MenuInfo)
			menuApi.POST("group", v1.Auth.SetMenuGroup)
		}

		userApi := api.Group("user")
		{
			userApi.GET("self", v1.User.Self)
			userApi.GET("list", v1.User.List)
			userApi.GET("info", v1.User.Info)
			userApi.POST("self", v1.User.ChangeSelf)
			userApi.POST("add", v1.User.Add)
			userApi.POST("update", v1.User.Change)
			userApi.POST("status", v1.User.ChangeStatus)
			userApi.POST("group", v1.User.SetUserGroup)
		}
	}

}
