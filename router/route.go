package router

import (
	"github.com/hideaki10/coredemo/framework"
	"github.com/hideaki10/coredemo/middleware"
)

func RegisterRouter(core *framework.Core) {
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}

}
