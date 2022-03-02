package router

import "github.com/hideaki10/coredemo/framework"

func SubjectDelController(c *framework.Context) error {
	c.Json(200, "delete success")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.Json(200, "list success")
	return nil
}

func SubjectAddController(c *framework.Context) error {
	c.Json(200, "add success")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.Json(200, "name success")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.Json(200, "get success")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.Json(200, "update success")
	return nil
}
