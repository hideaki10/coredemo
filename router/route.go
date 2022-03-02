package router

import (
	"time"

	"github.com/hideaki10/coredemo/framework"
)

func RegisterRouter(core *framework.Core) {
	core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second))

	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}

}
