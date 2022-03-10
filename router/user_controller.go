package router

import "github.com/hideaki10/coredemo/framework"

func UserLoginController(c *framework.Context) error {
	//c.Json(200, "login success")
	c.SetOkStatus().Json("login success")
	return nil
}
