package routers
import (
    "github.com/gin-gonic/gin"
    "github.com/windywork/bookrack/app/middleware"
    siteController "github.com/windywork/bookrack/app/modules/site/controllers"
    backendController "github.com/windywork/bookrack/app/modules/backend/controllers"
)

func InitRoutes(engine *gin.RouterGroup, appCtrl middleware.BaseController) {
    siteRouter := engine.Group("")
    {
      defCtrl := siteController.DefaultController{BaseController: appCtrl}
      siteRouter.GET("/", defCtrl.Index)
    }


    backendRouter := engine.Group("/backend")
    {
      defCtrl := backendController.DefaultController{BaseController: appCtrl}
      backendRouter.GET("/", defCtrl.Index)
      backendRouter.GET("/login", defCtrl.Login)
    }
}
