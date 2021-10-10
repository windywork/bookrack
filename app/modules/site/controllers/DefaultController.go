package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/unrolled/render"
    "net/http"
    "github.com/windywork/bookrack/app/middleware"
)

type DefaultController struct {
    middleware.BaseController
}

func (ctrl DefaultController) Index(c *gin.Context) {
    ctrl.Render.HTML(c.Writer, http.StatusOK, "site/default/index", gin.H{
        "Title": "Hello111111",
    }, render.HTMLOptions{Layout: "site/_layouts/main"})
}
