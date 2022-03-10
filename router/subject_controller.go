package router

import "github.com/hideaki10/coredemo/framework"

func SubjectDelController(c *framework.Context) error {
	//c.Json(200, "delete success")
	c.SetOkStatus().Json("delete success")
	return nil
}

func SubjectListController(c *framework.Context) error {
	//c.Json(200, "list success")
	c.SetOkStatus().Json("list success")
	return nil
}

func SubjectAddController(c *framework.Context) error {
	//c.Json(200, "add success")
	c.SetOkStatus().Json("add success")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	//c.Json(200, "name success")
	c.SetOkStatus().Json("name success")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	//c.Json(200, "get success")
	c.SetOkStatus().Json("get success")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	//c.Json(200, "update success")
	c.SetOkStatus().Json("update success")
	return nil
}
