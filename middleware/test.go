package middleware

import (
	"fmt"

	"github.com/hideaki10/coredemo/framework"
)

func Test1() framework.Controller {

	return func(c *framework.Context) error {
		fmt.Println("middleware.Test1")
		c.Next()
		fmt.Println("middleware post Test1")
		return nil
	}
}

func Test2() framework.Controller {

	return func(c *framework.Context) error {
		fmt.Println("middleware.Test2")
		c.Next()
		fmt.Println("middleware post Test2")
		return nil
	}
}

func Test3() framework.Controller {

	return func(c *framework.Context) error {
		fmt.Println("middleware.Test3")
		c.Next()
		fmt.Println("middleware post Test3")
		return nil
	}
}
